package controllers

import "net/http"

type MiddleWare struct {
}

func (m *MiddleWare) StateForOAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
