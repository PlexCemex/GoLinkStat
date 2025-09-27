package middleware

import (
	"context"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/pkg/jwt"
	"strings"
)
type key string

const(
	ContextEmailKey key = "ContextEmailKey"
)

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, data := jwt.NewJWT(conf.Auth.Secret).Parse(token)
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		newR := r.WithContext(ctx)
		next.ServeHTTP(w, newR)
	})
}
