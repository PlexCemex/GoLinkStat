package main

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":7080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 7080")
	server.ListenAndServe()
}

// 7.3
