package routes

import (
	"github.com/anopsy/surfAppv2/pkg/controllers"
	"github.com/gorilla/mux"
)

var SurfRoutes = func(router *mux.Router) {
	router.HandleFunc("/showResults/", controllers.ShowResults).Methods("GET")
	router.HandleFunc("/spots/", controllers.ShowSpots).Methods("GET")
	router.HandleFunc("/spots/", controllers.CreateSpot).Methods("POST")
	router.HandleFunc("/spots/{spotId}", controllers.UpdateSpot).Methods("PUT")
	router.HandleFunc("/spots/{spotId}", controllers.DeleteSpot).Methods("DELETE")
	router.HandleFunc("/spots/{spotId}", controllers.GetSpotByID).Methods("GET")
}
