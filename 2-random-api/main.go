package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/random", random)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Couldn't start server", err)
		panic(err)
	}
}

func random(w http.ResponseWriter, r *http.Request) {
	randomValue := rand.IntN(6) + 1
	fmt.Println("Random:", randomValue)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(strconv.Itoa(randomValue)))
}
