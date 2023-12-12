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

func Cors(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
