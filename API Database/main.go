package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

var stations = []station{
	{Code: "AFS", Name: "Ashford (Surrey)", Operator: "SWR", Platforms: 2},
	{Code: "STN", Name: "Staines", Operator: "SWR", Platforms: 2},
	{Code: "EGH", Name: "Egham", Operator: "SWR", Platforms: 2},
	{Code: "CLP", Name: "Clapham Jumction", Operator: "SWR", Platforms: 17},
}

func main() {
    router := gin.Default()
    router.GET("/stations", getStations)
	router.GET("/stations/:code", getStationByID)

    router.Run("localhost:8080")
}

// getStations responds with the list of all stations as JSON.
func getStations(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, stations)
}

func getStationByID(c *gin.Context) {
	code := c.Param("code")

	for _, s := range stations {
		if s.Code == code {
			c.IndentedJSON(http.StatusOK, s)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Station not found"})
}
