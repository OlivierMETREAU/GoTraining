package main

import (
	"net/http"

	"example.com/day02-http-server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/metrics"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Collect metrics for incoming HTTP requests automatically.
	r.Use(metrics.Collector(metrics.CollectorOpts{
		Host:  false,
		Proto: true,
		Skip: func(r *http.Request) bool {
			return r.Method != "OPTIONS"
		},
	}))
	r.Handle("/metrics", metrics.Handler())
	// Collect metrics for outgoing HTTP requests automatically.
	transport := metrics.Transport(metrics.TransportOpts{
		Host: true,
	})
	http.DefaultClient.Transport = transport(http.DefaultTransport)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Get("/hello", handler.SayHello)
	r.Get("/status", handler.GetStatus)
	r.Get("/date", handler.GetDate)
	http.ListenAndServe(":8080", r)
}
