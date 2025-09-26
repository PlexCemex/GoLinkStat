package middleware

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/jwt"
	"strings"
)

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(conf.Auth.Secret).Parse(token)
		fmt.Println(token)
		fmt.Println(isValid)
		fmt.Println(data)
		next.ServeHTTP(w, r)
	})
}
