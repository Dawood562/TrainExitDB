package main

import (
	. "trainAPI/structs"
	"trainAPI/database"
    "net/http"
    "github.com/gin-gonic/gin"
	"fmt"
)

// Will delete soon
var stations = []Station{
	{Code: "AFS", Name: "Ashford (Surrey)", Operator: "SWR", Platforms: 2},
	{Code: "STN", Name: "Staines", Operator: "SWR", Platforms: 2},
	{Code: "EGH", Name: "Egham", Operator: "SWR", Platforms: 2},
	{Code: "CLP", Name: "Clapham Jumction", Operator: "SWR", Platforms: 17},
}

func main() {
	// Initialise DB connection
	if err := database.GetConnection(); err != nil {
        fmt.Println("Failed to connect to DB: %w", err)
    }

    defer database.DB.Close()

    router := gin.Default()
    router.GET("/stations", getStations)
	router.GET("/stations/:code", getStationByID)

    router.Run("localhost:8080")
}

// getStations responds with the list of all stations as JSON.
func getStations(c *gin.Context) {
    stationss, err := database.GetStations()
    if err != nil {
         c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.IndentedJSON(http.StatusOK, stationss)
}

// getStations responds with information about one station when provided with a code
func getStationByID(c *gin.Context) {
	code := c.Param("code")

	station, err := database.GetStationByID(code);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Station not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, station)
}
