package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const api = "https://api.iextrading.com/1.0"

func main() {
	router := gin.Default()

	// Serve react build files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			data, err := chart(&http.Client{
				Timeout: 5 * time.Second,
			}, "meli", "1m")
			if err != nil {
				 c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}

			c.JSON(http.StatusOK, data)

		})
	}

	// start and run the server
	router.Run(":3000")
}

// Represents chart's json response from iextrading
type Chart struct {
	Date             string
	Open             float64
	Close            float64
	UnadjustedVolume int
	Change           float64
	ChangePercent    float64
	VWAP             float64
	High             float64
	Low              float64
	Volume           int
	Label            string
	ChangeOverTime   float64
}

// chart gives back chart data for a symbol in a range.
// Range must be either 1y or 1m
func chart(client *http.Client, symbol string, dataRange string) ([]*Chart, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be nil")
	}

	if dataRange == "" {
		dataRange = "1m"
	}
	resp, err := client.Get(fmt.Sprintf("%s/stock/%s/chart/%s", api, symbol, dataRange))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %v", resp.Status, string(body))
	}

	var result []*Chart
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err

}

// https://api.iextrading.com/1.0/stock/aapl/chart/1m
