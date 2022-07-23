package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type exchange struct {
	Code     string  `json:"code,omitempty"`
	Name     string  `json:"name"`
	MakerFee float64 `json:"maker_fee"`
	TakerFee float64 `json:"taker_fee"`
	ApiKey   string  `json:"api_key"`
	Secret   string  `json:"secret"`
	Test     bool    `json:"test"`
}

var exchanges = []exchange{
	{Code: "binance", Name: "Binance"},
}

func main() {
	router := gin.Default()
	router.GET("/api/exchange/v1/exchanges", getExchanges)
	router.GET("/api/exchange/v1/exchanges/:id", getExchangeByID)
	router.POST("/api/exchange/v1/exchanges", addExchange)
	router.Run("localhost:8080")
}

func getExchanges(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []exchange{})
}

func addExchange(c *gin.Context) {
	var exc exchange

	if err := c.BindJSON(&exc); err != nil {
		return
	}

	// Add the new exchange to the database

	c.IndentedJSON(http.StatusCreated, exc)
}

func getExchangeByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range exchanges {
		if a.Code == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "exchange not found"})
}
