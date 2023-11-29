package database

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	// Migrate the schema
	DB.AutoMigrate(&model.Elevator{}, &model.Hotel{}, &model.TravelHistory{})
}

func Close() {
	DB.Close()
}
