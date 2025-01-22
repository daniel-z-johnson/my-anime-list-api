package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/daniel-z-johnson/my-anime-list-api/rand"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("MAL api start")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		s, _ := rand.State()
		fmt.Fprintf(w, "hello %s %d\n", s, len(s))
	})
	logger.Info("Starting service on :1777")
	http.ListenAndServe(":1777", r)
}
