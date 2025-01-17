package repository

import (
	"flight/models"

	"gorm.io/gorm"
)

type FlightRepository struct {
	DB *gorm.DB
}

func (r *FlightRepository) GetStation() (*[]string, error) {
	var stations []string

	query := `
		(SELECT DISTINCT source AS station FROM flights)
		UNION
		(SELECT DISTINCT destination AS station FROM flights)
	`
	result := r.DB.Raw(query).Pluck("station", &stations)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stations, nil
}

func (r *FlightRepository) GetDirectFlights(source string, destination string, date string) ([]models.Flight, error) {
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

// func (r *FlightRepository) GetFlightsById(id int) (models.Flight, error) {
// 	var flight models.Flight

// 	result,err:= r.DB.First(&flight,id);
// 	if err!=nil{
// 		return nil,err
// 	}

// }
