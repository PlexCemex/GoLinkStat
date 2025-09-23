package main

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/link"
	"projects/GoLinkStat/internal/user"
	"projects/GoLinkStat/pkg/db"
	"projects/GoLinkStat/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	dataBase := db.NewDb(conf)
	router := http.NewServeMux()

	// Repository
	linkRepository := link.NewLinkRepository(dataBase)
	userRepository := user.NewUserRepository(dataBase)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	chain := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":7080",
		Handler: chain(router),
	}
	fmt.Println("Server is listening on port:", server.Addr)
	server.ListenAndServe()
}
