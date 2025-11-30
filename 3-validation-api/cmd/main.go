package main

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/internal/verify"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening on port 8081")
	server.ListenAndServe()
}
