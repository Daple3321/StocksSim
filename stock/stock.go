package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Stock struct {
	Ticker       string  `json:"ticker"`
	Amount       int     `json:"amount"`
	OriginalCost float64 `json:"originalCost"`
}

type StockInfo struct {
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
	Updated  int64   `json:"updated"`
	Currency string  `json:"currency"`
}

func (s StockInfo) String() string {
	return fmt.Sprintf("%s %f", s.Ticker, s.Price)
}

func FetchStockInfo(ticker string) (*StockInfo, error) {

	if ticker == "" {
		return nil, errors.New("can't get stock. no ticker name provided")
	}

	client := &http.Client{}
	apiRequest := fmt.Sprintf("https://api.api-ninjas.com/v1/stockprice?ticker=%s", ticker)

	apiKey := os.Getenv("ApiNinjas_API_KEY")
	if apiKey == "" {
		return nil, errors.New("ApiNinjas_API_KEY is missing in environment variables")
	}

	req, _ := http.NewRequest("GET", apiRequest, nil)
	req.Header.Add("X-Api-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stock StockInfo
	jsonErr := json.Unmarshal(body, &stock)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &stock, nil
}

func (s *Stock) GetStockGrowth() float64 {

	var growth float64

	stockInfo, fetchErr := FetchStockInfo(s.Ticker)
	if fetchErr != nil {
		fmt.Printf("Error fetching stock info: %s\n", fetchErr)
	}

	currentMarketValue := float64(s.Amount) * stockInfo.Price
	//fmt.Printf("Current makert val: %f\n", currentMarketValue)

	growth = (currentMarketValue / s.OriginalCost) - 1
	//fmt.Printf("Growth: %f\n", growth)

	return growth

}
