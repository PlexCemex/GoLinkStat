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
	return func(resWriter http.ResponseWriter, req *http.Request) {
		requestRegister, err := request.HandleBody[RegisterRequest](&resWriter, req)
		if err != nil {
			return
		}
		fmt.Println(requestRegister)
		res := RegisterResponse{
			Token: auth.Auth.Secret,
		}
		response.Json(res, resWriter, 201)
	}
}

func (auth *authHandler) Login() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		requestLogin, err := request.HandleBody[LoginRequest](&resWriter, req)
		if err != nil {
			return
		}
		fmt.Println(requestLogin)
		res := LoginResponse{
			Token: auth.Auth.Secret,
		}
		response.Json(res, resWriter, 200)
	}
}
