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
	PassengerID *int // Allow NULL initially
	Passenger   Passenger
	PNRnumber string
	FlightID    int `gorm:"not null"`
	Flight      Flight
	TicketType  string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}

type Passenger struct {
	gorm.Model
	BookingId  int
	Title      string
	FirstName  string
	LastName   string
	SeatNumber int
}
