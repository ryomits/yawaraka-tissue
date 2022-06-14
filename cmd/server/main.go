package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yawaraka-tissue/controller/build"
	"yawaraka-tissue/domain/problem"

	"github.com/go-chi/chi/v5"
)

func main() {
	s := Server()
	s.Addr = "0.0.0.0:8080"
	log.Fatal(s.ListenAndServe())
}

func Server() *http.Server {
	c := chi.NewRouter()
	c.NotFound(notfound)
	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Yawaraka Tissue"))
	})

	return &http.Server{
		Handler: c,
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	p := problem.NewBadRequest(problem.TypeBadRequest)

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(build.Error(p))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	p := problem.NewNotFound(
		problem.TypeResourceNotFound,
	).WithDetail(fmt.Sprintf("%s is not defined resource", r.URL.Path))

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(build.Error(p))
}
