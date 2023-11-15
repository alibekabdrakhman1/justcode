package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	go func() {
		debugRouter := chi.NewRouter()
		debugRouter.Mount("/debug", middleware.Profiler())

		debugServer := http.Server{
			Addr:    ":8081",
			Handler: debugRouter,
		}

		if err := debugServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	mainServer := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := mainServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
