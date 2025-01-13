package models

import "time"

type FlightResponse struct {
	ID            uint
	Source        string
	Destination   string
	Date          time.Time
	DepartureTime time.Time
	ArrivalTime   time.Time
	PriceSaver    float64
	PriceFlexi    float64
	Airline       string

	Layover            string  `json:"layover,omitempty"`
	CombinedPriceSaver float64 `json:"combined_price_saver,omitempty"`
	CombinedPriceFlexi float64 `json:"combined_price_flexi,omitempty"`
}
