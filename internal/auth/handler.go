package auth

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type authHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (auth *authHandler) Register() http.HandlerFunc {
	return func(resWritter http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}

func (auth *authHandler) Login() http.HandlerFunc {
	return func(resWritter http.ResponseWriter, req *http.Request) {
		fmt.Println(auth.Config.Auth.Secret)
		fmt.Println("Login")
		res := LoginResponse {
			Token: auth.Auth.Secret,
		}
		response.ResponseJson(res, resWritter, 200)
	}
}
