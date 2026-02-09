package main

import (
	"net/http"

	"example.com/day02-http-server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Get("/hello", handler.SayHello)
	r.Get("/status", handler.GetStatus)
	r.Get("/date", handler.GetDate)
	http.ListenAndServe(":8080", r)
}
