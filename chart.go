package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)

// Chart Represents chart's json response from iextrading
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