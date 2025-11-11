package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	. "trainAPI/structs"
)

var DB *sql.DB

// Connect to DB
func GetConnection() error {
	const path = "./database/flat.db"

	var err error
	DB, err = sql.Open("sqlite", path)

	// Check for errors when instantiating database connection 
	if err != nil {
		return fmt.Errorf("Failed to open DB: %w", err)
	}
	if err := DB.Ping(); err != nil {
		fmt.Errorf("Failed to ping DB: %w", err)
	}

	return err
}

// Get all stations
func GetStations() ([]Station, error) {
	if DB == nil {
        return nil, fmt.Errorf("Database not connected")
    }

	rows, err := DB.Query("SELECT Code, Name, Operator, Platforms FROM Station")

	// If error when querying
	if err != nil {
        return nil, err
    }
	defer rows.Close()

	// Create array of Station objects if rows are found
	var stations = []Station{}
    for rows.Next() {
        var s Station
        if err := rows.Scan(&s.Code, &s.Name, &s.Operator, &s.Platforms); err != nil {
            return nil, err
        }
        stations = append(stations, s)
    }

	// Make sure no error while iterating
	if err := rows.Err(); err != nil {
        return nil, err
    }

	return stations, nil
}

func GetStationByID(code string) (Station, error) {
	var station Station

	if DB == nil {
        return station, fmt.Errorf("Database not connected")
    }

	err := DB.QueryRow("SELECT Code, Name, Operator, Platforms FROM Station WHERE Code = ?", code).
		Scan(&station.Code, &station.Name, &station.Operator, &station.Platforms)

	if err == sql.ErrNoRows {
		return station, fmt.Errorf("Station not found")
	} else if err != nil {
		return station, fmt.Errorf("Query error: %v", err)
	}

	return station, nil
}