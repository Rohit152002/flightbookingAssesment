package repository

import (
	"flight/models"

	"gorm.io/gorm"
)

type FlightRepository struct {
	DB *gorm.DB
}

func (r *FlightRepository) GetStation() (*[]string, error) {
	var station []string

	result := r.DB.Model(&models.Flight{}).Distinct("source").Pluck("source", &station)
	if result.Error != nil {
		return nil, result.Error
	}
	return &station, nil
}

func (r *FlightRepository) GetDirectFlights(source, destination string, date string) ([]models.Flight, error) {
	var flights []models.Flight

	result := r.DB.Where("source = ? AND destination = ? AND date = ?", source, destination, date).Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}

	return flights, nil
}

func (r *FlightRepository) GetConnectingFlights(source, date string) ([]models.Flight, error) {
	var flights []models.Flight
	result := r.DB.Where("source = ? AND date = ?", source, date).Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}
	return flights, nil
}
