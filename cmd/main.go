package main

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/link"
	"projects/GoLinkStat/internal/stat"
	"projects/GoLinkStat/internal/user"
	"projects/GoLinkStat/pkg/db"
	"projects/GoLinkStat/pkg/event"
	"projects/GoLinkStat/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	dataBase := db.NewDb(conf)
	router := http.NewServeMux()
	eventbus := event.NewEventBus()

	// Repository
	linkRepository := link.NewLinkRepository(dataBase)
	userRepository := user.NewUserRepository(dataBase)
	statRepository := stat.NewStatRepository(dataBase)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventbus,
		StatRepository: statRepository,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventbus,
		Config:         conf,
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

	go statService.AddClick()

	fmt.Println("Server is listening on port:", server.Addr)
	server.ListenAndServe()
}
