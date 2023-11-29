package model

type TravelHistory struct {
	ID             uint `json:"id" gorm:"primary_key"`
	ElevatorID     uint `json:"elevator_id"`
	HotelID        uint `json:"hotel_id"`
	FloorsTraveled int  `json:"floors_traveled"`
}
