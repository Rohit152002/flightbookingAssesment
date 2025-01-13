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

// func (r *FlightServices) SearchFlights(source string, destination string, date string) (*[]models.FlightResponse, error) {
// 	if source == "" || destination == "" {
// 		return nil, errors.New("source and destination are required")
// 	}
// 	// if date.Compare(time.Now()) == -1 {
// 	// 	return nil, errors.New("The selected date is in the past. Please choose a valid future date.")
// 	// }

// 	return r.Repo.GetFlights(source, destination, date)
// }

func (s *FlightServices) GetFlights(source, destination string, date string) ([]models.FlightResponse, error) {

	directFlights, err := s.Repo.GetDirectFlights(source, destination, date)
	if err != nil {
		return nil, err
	}

	var directFlightResponses []models.FlightResponse
	for _, flight := range directFlights {
		directFlightResponses = append(directFlightResponses, models.FlightResponse{
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

	layoverFlights, err := s.Repo.GetConnectingFlights(source, date)
	if err != nil {
		return nil, err
	}

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

	allFlights := append(directFlightResponses, connectingFlightResponses...)

	return allFlights, nil
}
