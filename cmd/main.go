package main

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/internal/hello"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	server := http.Server{
		Addr:    ":7080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 7080")
	server.ListenAndServe()
}
