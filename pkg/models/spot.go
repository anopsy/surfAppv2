package models

import (
	"time"

	"github.com/anopsy/surfAppv2/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Spot struct {
	gorm.Model
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Weather struct {
	gorm.Model
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Lat         string    `json:"lat"`
	Long        string    `json:"long"`
	Time        time.Time `json:"time"`
	Swell       float64   `json:"swell"`
	SwellPeriod float64   `json:"period"`
	Wave        float64   `json:"wave"`
	Wind        float64   `json:"wind"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Spot{})

}

func (s *Spot) CreateSpot() *Spot {
	db.Create(&s)
	return s
}

func ShowSpots() []Spot {
	var Spots []Spot
	db.Find(&Spots)
	return Spots
}

func ShowResults() []Spot {
	var Spots []Spot
	//query db
	return Spots
}

func GetSpotByID(ID int64) (*Spot, *gorm.DB) {
	var getSpot Spot
	db := db.Where("ID=?", ID).Find(&getSpot)
	return &getSpot, db

}

func DeleteSpot(ID int64) Spot {
	var spot Spot
	db.Where("ID=?", ID).Delete(spot)
	return spot
}
