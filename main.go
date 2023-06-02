package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taher267/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStroreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5001", r))
}
