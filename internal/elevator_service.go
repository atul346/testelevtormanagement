package service

import (
	"github.com/atul346/elevator-management-go/internal/model"
	"github.com/atul346/elevator-management-go/internal/repository"
)

func GetElevators() ([]model.Elevator, error) {
	elevators, err := repository.GetElevators()
	if err != nil {
		return nil, err
	}
	return elevators, nil
}
