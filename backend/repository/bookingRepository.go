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
	if err := booking.DB.Preload("Flight").First(&newBooking, newBooking.ID).Error; err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &newBooking, nil
}

func (booking *BookRepository) FindByReferenceNumber(referenceNo string) (*models.Booking, error) {
	var booked models.Booking

	// Use Preload to fetch related Flight and Passengers in a single query
	if err := booking.DB.Preload("Flight").Preload("Passengers").
		Where("booking_reference_no = ?", referenceNo).First(&booked).Error; err != nil {
		return nil, err
	}

	return &booked, nil
}
