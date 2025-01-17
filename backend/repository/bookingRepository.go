package repository

import (
	"errors"
	"flight/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (booking *BookRepository) Create(newBooking models.Booking) (*models.Booking, error) {
	if newBooking.PassengerID == nil {
		return nil, errors.New("PassengerID cannot be null")
	}
	if newBooking.FlightID == 0 {
		return nil, errors.New("FlightID cannot be zero")
	}
	if newBooking.TicketType == "" {
		return nil, errors.New("TicketType cannot be empty")
	}
	if newBooking.Price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	result := booking.DB.Create(&newBooking)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newBooking, nil
}
