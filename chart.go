package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Chart Represents chart's json response from iextrading
type Chart struct {
	Date             string  `json:"date"`
	Open             float64 `json:"open"`
	Close            float64 `json:"close"`
	UnadjustedVolume int     `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	VWAP             float64 `json:"vwap"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Volume           int     `json:"volume"`
	Label            string  `json:"label"`
	ChangeOverTime   float64 `json:"changeOverTime"`
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
