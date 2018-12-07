package main

import (
	"sync"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func StockHandler(c *gin.Context) {

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	result := make([][]*Chart, 0)

	c.Header("Content-Type", "application/json")
	dataRange := c.DefaultQuery("range", "1m")
	symbols := c.Query("symbols")

	if symbols == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "symbols must be defined",
		})
		return
	}

	splittedSymbols := strings.Split(symbols, ",")

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, symbol := range splittedSymbols {
		if symbol != "" {
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				data, err := chart(&httpClient, s, dataRange)
				if err != nil {
					c.JSON(http.StatusForbidden, gin.H{
						"error": err.Error(),
					})
					return
				}

				mutex.Lock()
				result = append(result, data)
				mutex.Unlock()
			}(symbol)
		}
	}

	wg.Wait()
	c.JSON(http.StatusOK, result)
	
}