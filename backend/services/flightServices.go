package services

import (
	"flight/models"
	"flight/repository"
)

type FlightServices struct {
	Repo repository.FlightRepository
}

func (r *FlightServices) GetDistinctSources() (*[]string, error) {
	return r.Repo.GetStation()
}

func (s *FlightServices) mapFlightDetails(flights []models.Flight) []models.FlightResponse {
	var flightResponses []models.FlightResponse
	for _, flight := range flights {
		flightResponses = append(flightResponses, models.FlightResponse{
			ID:            flight.ID,
			Source:        flight.Source,
			Destination:   flight.Destination,
			Date:          flight.Date,
			DepartureTime: flight.DepartureTime,
			ArrivalTime:   flight.ArrivalTime,
			PriceSaver:    flight.PriceSaver,
			PriceFlexi:    flight.PriceFlexi,
			Airline:       flight.Airline,
		})
	}
	return flightResponses
}

func (s *FlightServices) mapConnectingFlights(layoverFlights []models.Flight, destination, date string) ([]models.FlightResponse, error) {
	var connectingFlightResponses []models.FlightResponse
	for _, layover := range layoverFlights {
		nextLegs, err := s.Repo.GetDirectFlights(layover.Destination, destination, date)
		if err != nil {
			return nil, err
		}
		for _, nextLeg := range nextLegs {
			connectingFlightResponses = append(connectingFlightResponses, models.FlightResponse{
				Source:             layover.Source,
				Destination:        nextLeg.Destination,
				Date:               layover.Date,
				DepartureTime:      layover.DepartureTime,
				ArrivalTime:        nextLeg.ArrivalTime,
				PriceSaver:         layover.PriceSaver + nextLeg.PriceSaver,
				PriceFlexi:         layover.PriceFlexi + nextLeg.PriceFlexi,
				Airline:            layover.Airline + " / " + nextLeg.Airline,
				Layover:            layover.Destination,
				CombinedPriceSaver: layover.PriceSaver + nextLeg.PriceSaver,
				CombinedPriceFlexi: layover.PriceFlexi + nextLeg.PriceFlexi,
			})
		}
	}
	return connectingFlightResponses, nil
}

func (s *FlightServices) GetFlights(source, destination, date string) ([]models.FlightResponse, error) {
	directFlights, err := s.Repo.GetDirectFlights(source, destination, date)
	if err != nil {
		return nil, err
	}
	directFlightResponses := s.mapFlightDetails(directFlights)

	layoverFlights, err := s.Repo.GetConnectingFlights(source, date)
	if err != nil {
		return nil, err
	}
	connectingFlightResponses, err := s.mapConnectingFlights(layoverFlights, destination, date)
	if err != nil {
		return nil, err
	}

	allFlights := append(directFlightResponses, connectingFlightResponses...)
	return allFlights, nil
}

func (s *FlightServices) GetFlightsWithReturn(source, destination, departureDate, returnDate string) ([]models.FlightResponse, []models.FlightResponse, error) {
	onwardFlights, err := s.GetFlights(source, destination, departureDate)
	if err != nil {
		return nil, nil, err
	}

	returnFlights, err := s.GetFlights(destination, source, returnDate)
	if err != nil {
		return nil, nil, err
	}

	// allFlights := append(onwardFlights, returnFlights...)
	return onwardFlights, returnFlights, nil
}
