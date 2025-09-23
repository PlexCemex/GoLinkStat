package auth

import (
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/jwt"
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

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		jwtToken, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: jwtToken,
		}
		response.Json(data, w, 200)
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogin, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		existedEmail, err := handler.AuthService.Login(requestLogin.Email, requestLogin.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		jwtToken, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(existedEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Token: jwtToken,
		}
		response.Json(data, w, 200)
	}
}
