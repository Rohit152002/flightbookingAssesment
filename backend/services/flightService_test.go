package services

import (
	"flight/repository"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s'", err)
	}

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("gormDatabase %s", err)
	}
	return gormDb, mock
}

func TestFlightServices(t *testing.T) {
	mockDb, mock := mockDatabase()
	flightRepository := repository.FlightRepository{DB: mockDb}
	flightServices := FlightServices{Repo: flightRepository}
	layout := "2006-01-02 15:04:05"
	departureTime, _ := time.Parse(layout, "2025-01-15 06:00:00")
	arrivalTime, _ := time.Parse(layout, "2025-01-15 08:00:00")
	layoverDepartureTime, _ := time.Parse(layout, "2025-01-15 06:00:00")
	layoverArrivalTime, _ := time.Parse(layout, "2025-01-15 08:00:00")
	// connectingArrivalTime, _ := time.Parse(layout, "2025-01-15 11:00:00")

	columns := []string{"id", "source", "destination", "date", "departureTime", "arrivalTime", "saverPrice", "flexiPrice", "airline"}
	t.Run("Get Stations", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"station"}).
			AddRow("delhi").AddRow("mumbai")

		mock.ExpectQuery(`(SELECT DISTINCT source AS station FROM flights)
			UNION
			(SELECT DISTINCT destination AS station FROM flights)`).
			WillReturnRows(rows)
		result, err := flightServices.GetDistinctSources()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		expectedStations := []string{"delhi", "mumbai"}
		assert.ElementsMatch(t, expectedStations, *result, "The returned stations do not match the expected ones.")
	})

	t.Run("GetStationsError", func(t *testing.T) {
		mock.ExpectQuery(`(SELECT DISTINCT source AS station FROM flights)
			UNION
			(SELECT DISTINCT destination AS station FROM flights)`).
			WillReturnError(assert.AnError)

		result, err := flightServices.GetDistinctSources()

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetFlights_Success", func(t *testing.T) {
		// Mocking the direct flights query
		directRows := sqlmock.NewRows(columns).
			AddRow(1, "delhi", "mumbai", departureTime, departureTime, arrivalTime, 100, 150, "AirIndia").
			AddRow(2, "delhi", "mumbai", departureTime, departureTime, arrivalTime, 120, 180, "Indigo")

		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND destination = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "mumbai", "2025-01-15").
			WillReturnRows(directRows)

		// Mocking the connecting flights query
		connectingRows := sqlmock.NewRows(columns).
			AddRow(3, "delhi", "bangalore", layoverDepartureTime, layoverDepartureTime, layoverArrivalTime, 50, 80, "AirIndia").
			AddRow(4, "bangalore", "mumbai", layoverDepartureTime, layoverDepartureTime, layoverArrivalTime, 60, 90, "Indigo")

		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "2025-01-15").
			WillReturnRows(connectingRows)

		// Mocking GetDirectFlights within mapConnectingFlights
		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND destination = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").WithArgs("bangalore", "mumbai", "2025-01-15").
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(5, "bangalore", "mumbai", time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 15, 8, 30, 0, 0, time.UTC), time.Date(2025, 1, 15, 11, 0, 0, 0, time.UTC), 70, 100, "SpiceJet"))

		// Call the GetFlights method
		flights, err := flightServices.GetFlights("delhi", "mumbai", "2025-01-15")

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, flights, 3)

		// Validate direct flights responses
		assert.Equal(t, uint(1), flights[0].ID)
		assert.Equal(t, "delhi", flights[0].Source)
		assert.Equal(t, "mumbai", flights[0].Destination)
		assert.Equal(t, "AirIndia", flights[0].Airline)

		// Validate connecting flights responses
		assert.Equal(t, "delhi", flights[2].Source)
		assert.Equal(t, "mumbai", flights[2].Destination)
		assert.Equal(t, "AirIndia / SpiceJet", flights[2].Airline)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetFlights_Error", func(t *testing.T) {
		// Simulate an error when getting direct flights
		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND destination = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "mumbai", "2025-01-15").
			WillReturnError(assert.AnError)

		// Call the GetFlights method
		flights, err := flightServices.GetFlights("delhi", "mumbai", "2025-01-15")

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, flights)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
