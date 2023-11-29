package repository

import (
    "github.com/atul346/elevator-management-go/internal/model"
    "github.com/atul346/elevator-management-go/pkg/database"
)

func GetElevators() ([]model.Elevator, error) {
    var elevators []model.Elevator
    if err := database.DB.Find(&elevators).Error; err != nil {
        return nil, err
    }
    return elevators, nil
}




