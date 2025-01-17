package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	Source        string
	Destination   string
	Date          time.Time
	DepartureTime time.Time
	ArrivalTime   time.Time
	PriceSaver    float64
	PriceFlexi    float64
	Airline       string
}

type Booking struct {
	gorm.Model
	PassengerId int
	Passenger   Passenger
	FlightId    int
	Flight      Flight
	FlightType  string
	Price       float64
}

type Passenger struct {
	gorm.Model
	Title     string
	FirstName string
	LastName  string
	Seat      int
}
