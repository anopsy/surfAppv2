package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anopsy/surfAppv2/pkg/models"
	"github.com/anopsy/surfAppv2/pkg/utils"
	"github.com/gorilla/mux"
)

var NewSpot models.Spot

func ShowSpots(w http.ResponseWriter, r *http.Request) {
	newSpots := models.ShowResults()
	res, _ := json.Marshal(newSpots)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowResults(w http.ResponseWriter, r *http.Request) {
	newResults := models.ShowResults()
	res, _ := json.Marshal(newResults)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSpotByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotId := vars["spotId"]
	ID, err := strconv.ParseInt(spotId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	spotDetails, _ := models.GetSpotByID(ID)
	res, _ := json.Marshal(spotDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateSpot(w http.ResponseWriter, r *http.Request) {
	CreateSpot := &models.Spot{}
	utils.ParseBody(r, CreateSpot)
	s := CreateSpot.CreateSpot()
	res, _ := json.Marshal(s)
	//w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotId := vars["spotId"]
	ID, err := strconv.ParseInt(spotId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	spot := models.DeleteSpot(ID)
	res, _ := json.Marshal(spot)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	var updateSpot = &models.Spot{}
	utils.ParseBody(r, updateSpot)
	vars := mux.Vars(r)
	spotID := vars["spotId"]
	ID, err := strconv.ParseInt(spotID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	spotDetails, db := models.GetSpotByID(ID)
	if updateSpot.Name != "" {
		spotDetails.Name = updateSpot.Name
	}
	if updateSpot.Lat != "" {
		spotDetails.Lat = updateSpot.Lat
	}

	if updateSpot.Long != "" {
		spotDetails.Long = updateSpot.Long
	}

	if updateSpot.Conditions != nil {
		spotDetails.Conditions = updateSpot.Conditions
	}

	db.Save(&spotDetails)
	res, _ := json.Marshal(spotDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
