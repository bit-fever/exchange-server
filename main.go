package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	router := gin.Default()
	router.GET("/api/exchange/v1/exchanges", getExchanges)
	router.Run("localhost:8080")
}

func getExchanges(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []exchange{})
}

