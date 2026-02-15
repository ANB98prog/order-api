package main

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/handler"
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/repository"
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/service"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/cache"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/db"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/jwt"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/logging"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/middlewares"
	"log"
	"net/http"
)

func main() {
	configs.ReadEnvironmentVariables()
	conf := configs.LoadConfig()
	logConfig, err := logging.ReadLogConfig()
	if err != nil {
		panic(err)
	}
	logging.ConfigureLogrus(logConfig)

	router := http.NewServeMux()

	// Databases
	dbConnection := db.NewDb(&conf.Db)
	redisClient := cache.NewRedisClient(&conf.Cache)

	// Repositories
	authCodeRepo := repository.NewRedisAuthCodeRepository(redisClient)
	userRepo := repository.NewUserRepository(dbConnection)
	productRepo := repository.NewProductRepository(dbConnection)
	orderRepo := repository.NewOrderRepository(dbConnection)

	// Services
	authCodeService := service.NewAuthCodeService(authCodeRepo)
	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(orderRepo, productRepo, userRepo)

	// JWT
	j := jwt.NewJWT(conf.Auth.Secret)

	// Handlers
	authHandlerDeps := handler.AuthHandlerDeps{
		UserService:     userService,
		AuthCodeService: authCodeService,
		JWT:             j,
	}
	orderHandlerDeps := handler.OrderHandlerDeps{
		OrderService: orderService,
	}

	handler.NewAuthHandler(router, authHandlerDeps)
	handler.NewOrderHandler(router, orderHandlerDeps, conf)

	stack := middlewares.Chain(middlewares.Logging)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Listening on port 8081")
	log.Fatal(server.ListenAndServe())
}
