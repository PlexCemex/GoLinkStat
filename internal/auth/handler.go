package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/response"
	"regexp"
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
		var requestLogin LoginRequest
		err := json.NewDecoder(req.Body).Decode(&requestLogin)
		if err != nil {
			response.ResponseJson(err.Error(), resWritter, 402)
			return
		}
		if requestLogin.Email == "" {
			response.ResponseJson("Empty or wrong email", resWritter, 402)
			return
		}
		regEmail, _ := regexp.Compile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)
		if !regEmail.MatchString(requestLogin.Email) {
			response.ResponseJson("Wrong email", resWritter, 402)
			return
		}
		if requestLogin.Password == "" {
			response.ResponseJson("Empty or wrong password", resWritter, 402)
			return
		}
		fmt.Println(requestLogin)
		res := LoginResponse{
			Token: auth.Auth.Secret,
		}
		response.ResponseJson(res, resWritter, 200)
	}
}
