package repository

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/atul346/elevator-management-go/pkg/database"
)

func GetHotels() ([]model.Hotel, error) {
	var hotels []model.Hotel
	if err := database.DB.Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}
