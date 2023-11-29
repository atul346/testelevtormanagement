package repository

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/atul346/elevator-management-go/pkg/database"
)

func GetTravelHistories() ([]model.TravelHistory, error) {
	var travelHistories []model.TravelHistory
	if err := database.DB.Find(&travelHistories).Error; err != nil {
		return nil, err
	}
	return travelHistories, nil
}
