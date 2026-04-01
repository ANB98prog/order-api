package orderapi

import (
	"context"
	"errors"
	"github.com/ANB98prog/order-api/internal/config"
	"github.com/ANB98prog/order-api/internal/handler"
	"github.com/ANB98prog/order-api/internal/repository"
	"github.com/ANB98prog/order-api/internal/service"
	"github.com/ANB98prog/order-api/pkg/db"
	"github.com/ANB98prog/order-api/pkg/jwt"
	"github.com/ANB98prog/order-api/pkg/middlewares"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"time"
)

const defaultHTTPAddr = ":8081"

type Dependencies struct {
	Logger      *slog.Logger
	Config      config.Config
	DB          *db.Db
	RedisClient *redis.Client
}

type Server struct {
	logger     *slog.Logger
	cfg        config.Config
	httpServer *http.Server
}

func New(deps Dependencies) *Server {
	router := http.NewServeMux()

	// Repositories
	authCodeRepo := repository.NewRedisAuthCodeRepository(deps.RedisClient)
	userRepo := repository.NewUserRepository(deps.DB)
	productRepo := repository.NewProductRepository(deps.DB)
	orderRepo := repository.NewOrderRepository(deps.DB)

	// Services
	authCodeService := service.NewAuthCodeService(authCodeRepo)
	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(orderRepo, productRepo, userRepo)
	productService := service.NewProductService(productRepo)

	// JWT
	tokenManager := jwt.NewJWT(deps.Config.Auth.Secret)
	authorization := middlewares.Authorization(tokenManager)

	// Handlers
	handler.NewAuthHandler(router, handler.AuthHandlerDeps{
		UserService:     userService,
		AuthCodeService: authCodeService,
		JWT:             tokenManager,
	})
	handler.NewOrderHandler(router, handler.OrderHandlerDeps{
		OrderService:  orderService,
		Authorization: authorization,
	})
	handler.NewProductHandler(router, handler.ProductHandlerDeps{
		ProductService: productService,
		Authorization:  authorization,
	})

	stack := middlewares.Chain(middlewares.Logging)

	httpAddr := deps.Config.Addr
	if httpAddr == "" {
		httpAddr = defaultHTTPAddr
	}

	return &Server{
		logger: deps.Logger,
		cfg:    deps.Config,
		httpServer: &http.Server{
			Addr:    defaultHTTPAddr,
			Handler: stack(router),
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting order api server", slog.String("addr", s.httpServer.Addr))

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Stop(shutdownCtx); err != nil {
			s.logger.Error("shutdown server", slog.String("error", err.Error()))
		}
	}()

	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
