package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"tideWatchAPI/utils"
)

func CreateConnection() *gorm.DB {
	// TODO: Set this up in env variables later
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=tidewatch dbname=tidewatch password=abc123 sslmode=disable")

	utils.ErrorCheck(err)

	log.Println("Connection successful")
	return db
}

func GetAllStationsByType(db *gorm.DB, v interface{}) bool {
	if db.Find(v).RecordNotFound() {
		return false
	}

	return true
}

func GetAllStationsById(db *gorm.DB, v interface{}, id string) bool {
	if db.Limit(20).Where("id LIKE ?", "%"+id+"%").Find(v).RecordNotFound() {
		return false
	}

	return true
}

func GetAllStationsByName(db *gorm.DB, v interface{}, name string) bool {

	if db.Limit(20).Where("name LIKE ?", "%"+name+"%").Find(v).RecordNotFound() {
		return false
	}

	return true
}
