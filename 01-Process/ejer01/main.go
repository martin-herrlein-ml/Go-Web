package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	//server
	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong!"))
	})

	//run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
