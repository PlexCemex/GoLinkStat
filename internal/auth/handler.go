package auth

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/request"
	"projects/GoLinkStat/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}
type authHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (auth *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		auth.AuthService.Register(body.Email, body.Password, body.Name)
	}
}

func (auth *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogin, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		existedEmail, err := auth.AuthService.Login(requestLogin.Email, requestLogin.Password)
		if err != nil {
			return
		}
		fmt.Println(existedEmail, err)
		data := LoginResponse{
			Token: auth.Auth.Secret,
		}
		response.Json(data, w, 200)
	}
}
