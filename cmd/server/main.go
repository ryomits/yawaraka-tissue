package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	s := Server()
	s.Addr = "0.0.0.0:8080"
	log.Fatal(s.ListenAndServe())
}

func Server() *http.Server {
	c := chi.NewRouter()
	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Yawaraka Tissue"))
	})

	return &http.Server{
		Handler: c,
	}
}
