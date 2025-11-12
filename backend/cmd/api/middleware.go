package main

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			// Specify that credentials are allowed - although with wildcard origin this has no effect
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			// Specify the allowed methods
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			// Specify the allowed headers
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
