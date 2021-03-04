package main

import (
	"log"
	"net/http"

	"../../pkg/routes"
	"github.com/gorilla/mux"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterQuestionRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
