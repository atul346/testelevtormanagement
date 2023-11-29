package service

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/atul346/elevator-management-go/internal/repository"
)

func GetTravelHistories() ([]model.TravelHistory, error) {
	travelHistories, err := repository.GetTravelHistories()
	if err != nil {
		return nil, err
	}
	return travelHistories, nil
}
