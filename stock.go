package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func StockHandler(c *gin.Context)  {
	c.Header("Content-Type", "application/json")
	dataRange := c.DefaultQuery("range", "1m")
	symbols := c.Query("symbols")
	
	if symbols == "" {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "symbols must be defined",
		})
	}

	splittedSymbols := strings.Split(symbols, ",")

	data1, err := chart(&http.Client{
		Timeout: 5 * time.Second,
	}, splittedSymbols[0], dataRange)

	data2, err := chart(&http.Client{
		Timeout: 5 * time.Second,
	}, splittedSymbols[1], dataRange)

	data3, err := chart(&http.Client{
		Timeout: 5 * time.Second,
	}, splittedSymbols[2], dataRange)


	if err != nil {
		c.JSON(http.StatusForbidden, gin.H {
			"error": err.Error(),
		})
		return
	}

	result := make([][]*Chart, 3)
	result[0] = data1
	result[1] = data2
	result[2] = data3

	c.JSON(http.StatusOK, result)
}