package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)
	server := http.Server{
		Addr: ":7080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 7080")
	server.ListenAndServe()
}