package player

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"gameroll.com/StocksSim/stock"
)

type StockFetcher interface {
	Fetch(ticker string) (*stock.StockInfo, error)
}

type DefaultStockFetcher struct{}

func (d *DefaultStockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

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

	//fmt.Printf("Response status: %s\n", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	sb := strings.Builder{}
	for {
		end := scanner.Scan()
		if !end {
			break
		}

		sb.Write([]byte(scanner.Text()))
	}

	var stock stock.StockInfo
	jsonErr := json.Unmarshal([]byte(sb.String()), &stock)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &stock, nil
}
