package main

import (
	"log"
	"net/http"

	"github.com/anopsy/surfAppv2/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.SurfRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
