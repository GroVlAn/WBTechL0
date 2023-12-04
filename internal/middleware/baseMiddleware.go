package middleware

import "net/http"

func SkipFavicon(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			return
		}

		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
