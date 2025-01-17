package repository

import (
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



func TestFlight(t *testing.T) {
	mockDb, mock := mockDatabase()
	flightRepository := FlightRepository{DB: mockDb}

	columns := []string{"id", "source", "destination", "date", "departureTime", "arrivalTime", "saverPrice", "flexiPrice", "airline"}
	t.Run("getStation", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"station"}).
			AddRow("delhi").AddRow("mumbai")

		mock.ExpectQuery(`(SELECT DISTINCT source AS station FROM flights)
			UNION
			(SELECT DISTINCT destination AS station FROM flights)`).
			WillReturnRows(rows)

		result, err := flightRepository.GetStation()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		expectedStations := []string{"delhi", "mumbai"}
		assert.ElementsMatch(t, expectedStations, *result, "The returned stations do not match the expected ones.")

	})
	t.Run("getStationError", func(t *testing.T) {
		mock.ExpectQuery(`(SELECT DISTINCT source AS station FROM flights)
			UNION
			(SELECT DISTINCT destination AS station FROM flights)`).
			WillReturnError(assert.AnError)

		result, err := flightRepository.GetStation()

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("searchFlight", func(t *testing.T) {

		rows := sqlmock.NewRows(columns).
			AddRow(1, "delhi", "mumbai", time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), 2000, 3000, "indigo")

		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND destination = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "mumbai", "2025-01-08").
			WillReturnRows(rows)

		result, err := flightRepository.GetDirectFlights("delhi", "mumbai", "2025-01-08")

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		for _, value := range result {
			assert.Equal(t, value.Source, "delhi")
			assert.Equal(t, value.Destination, "mumbai")
		}
	})

	t.Run("connectingFlights", func(t *testing.T) {

		rows := sqlmock.NewRows(columns).
			AddRow(1, "delhi", "mumbai", time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 8, 0, 0, 0, 0, time.UTC), 2000, 3000, "indigo")

		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "2025-01-08").
			WillReturnRows(rows)

		result, err := flightRepository.GetConnectingFlights("delhi", "2025-01-08")

		assert.NoError(t, mock.ExpectationsWereMet())
		assert.NoError(t, err)
		for _, value := range result {
			assert.Equal(t, value.Source, "delhi")
			assert.Equal(t, value.Destination, "mumbai")
		}
	})
	t.Run("SearchFlightError", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND destination = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "mumbai", "2025-01-08").
			WillReturnError(assert.AnError)

		result, err := flightRepository.GetDirectFlights("delhi", "mumbai", "2025-01-08")

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("ConnectingFlightError", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? AND date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "2025-01-08").
			WillReturnError(assert.AnError)

		result, err := flightRepository.GetConnectingFlights("delhi", "2025-01-08")
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})
}
