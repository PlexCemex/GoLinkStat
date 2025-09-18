package main

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/link"
	"projects/GoLinkStat/pkg/db"
	"projects/GoLinkStat/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	dataBase := db.NewDb(conf)
	router := http.NewServeMux()

	// Repository
	linkRepository := link.NewLinkRepository(dataBase)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	chain := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
		middleware.IsAuthed,
	)

	server := http.Server{
		Addr:    ":7080",
		Handler: chain(router),
	}
	fmt.Println("Server is listening on port:", server.Addr)
	server.ListenAndServe()
}

// 10.5
