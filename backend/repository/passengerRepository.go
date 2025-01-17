package repository

import (
	"errors"
	"flight/models"

	"gorm.io/gorm"
)

type PassengerRepository struct {
	*gorm.DB
}

func (repo *PassengerRepository) Create(newPassenger models.Passenger) (*models.Passenger, error) {
	if newPassenger.Title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	if newPassenger.FirstName == "" {
		return nil, errors.New("FirstName cannot be empty")
	}
	if newPassenger.LastName == "" {
		return nil, errors.New("LastName cannot be empty")
	}
	if newPassenger.SeatNumber <= 0 {
		return nil, errors.New("SeatNumber must be greater than zero")
	}

	result := repo.DB.Create(&newPassenger)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newPassenger, nil
}
