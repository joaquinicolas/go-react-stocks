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
		return
	}

	splittedSymbols := strings.Split(symbols, ",")

	httpClient := http.Client {
		Timeout: 5 * time.Second,
	}
	data1, err := chart(&httpClient, splittedSymbols[0], dataRange)
	var data2 []*Chart
	var data3 []*Chart

	if len(splittedSymbols) > 1{
		data2, err = chart(&httpClient, splittedSymbols[1], dataRange)
	}
	if len(splittedSymbols) > 2 {
		data3, err = chart(&httpClient, splittedSymbols[2], dataRange)
	}

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H {
			"error": err.Error(),
		})
		return
	}

	result := make([][]*Chart, 1)
	result[0] = data1

	if len(data2) > 0 {
		result = append(result, data2)
	}
	
	if len(data3) > 0 {
		result = append(result, data3)
	}
	

	c.JSON(http.StatusOK, result)
}