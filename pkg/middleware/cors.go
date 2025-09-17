package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", origin)
			header.Set("Access-Control-Allow-Credentials", "true")
			if r.Method == http.MethodOptions {
				header.Set("Access-Control-Allow-Methods", "Get,Put,Post,Delete,Head,Patch")
				header.Set("Access-Control-Allow-Headers", "authorization,content-type,content-lenght")
				header.Set("Access-Control-Max-Age", "86400")
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
