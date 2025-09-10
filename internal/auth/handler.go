package auth

import (
	"fmt"
	"net/http"
)

type authHandler struct {}

func NewAuthHandler(router *http.ServeMux) {
	handler := &authHandler{}
	router.HandleFunc("POST /auth/register",handler.Register())
	router.HandleFunc("POST /auth/login",handler.Login())
}

func (auth *authHandler) Register() http.HandlerFunc {
	return func(resWritter http.ResponseWriter, req *http.Request){
		fmt.Println("Register")
	} 

}

func (auth *authHandler) Login() http.HandlerFunc {
	return func(resWritter http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
	}
}
