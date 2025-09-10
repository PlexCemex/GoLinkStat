package main

import (
	"fmt"
	"net/http"
	// "projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
)

func main() {
	// conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router)
	server := http.Server{
		Addr:    ":7080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 7080")
	server.ListenAndServe()
}
// 6.6