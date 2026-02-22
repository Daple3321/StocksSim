package player

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gameroll.com/StocksSim/stock"
)

type StockFetcher interface {
	Fetch(ticker string) (*stock.StockInfo, error)
}

type DefaultStockFetcher struct {
	client *http.Client
}

func (d *DefaultStockFetcher) getClient() *http.Client {
	if d.client == nil {
		d.client = &http.Client{Timeout: 4 * time.Second}
	}
	return d.client
}

func (d *DefaultStockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

	if ticker == "" {
		return nil, errors.New("can't get stock. no ticker name provided")
	}

	apiRequest := fmt.Sprintf("https://api.api-ninjas.com/v1/stockprice?ticker=%s", ticker)

	apiKey := os.Getenv("ApiNinjas_API_KEY")
	if apiKey == "" {
		return nil, errors.New("ApiNinjas_API_KEY is missing in environment variables")
	}

	req, _ := http.NewRequest("GET", apiRequest, nil)
	req.Header.Add("X-Api-Key", apiKey)

	resp, err := d.getClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stock stock.StockInfo
	jsonErr := json.Unmarshal(body, &stock)
	if jsonErr != nil {
		return nil, jsonErr
	}

	//fmt.Println(stock)

	return &stock, nil
}
