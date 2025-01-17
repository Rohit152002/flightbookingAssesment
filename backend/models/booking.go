package models

type BookingDTO struct {
	FlightID           int
	PassengerFirstName string
	PassengerLastName  string
	PassengerEmail     string
	PriceSelect        int
}
