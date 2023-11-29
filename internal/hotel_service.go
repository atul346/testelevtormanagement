package service

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/atul346/elevator-management-go/internal/repository"
)

func GetHotels() ([]model.Hotel, error) {
	hotels, err := repository.GetHotels()
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
