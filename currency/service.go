package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type RateService interface {
	GetExchangeRate(base string, to string) (float64, error)
}

type DefaultRateService struct {
	client *http.Client
}

func (d *DefaultRateService) getClient() *http.Client {
	if d.client == nil {
		d.client = &http.Client{Timeout: 4 * time.Second}
	}
	return d.client
}

func (d *DefaultRateService) GetExchangeRate(base string, to string) (float64, error) {
	if base == "" {
		return 0, errors.New("no base currency provided")
	}
	if to == "" {
		return 0, errors.New("no target currency provided")
	}

	apiKey := os.Getenv("CurrencyExchange_API_KEY")
	if apiKey == "" {
		return 0, errors.New("CurrencyExchange_API_KEY is missing in environment variables")
	}
	apiRequest := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", apiKey, base, to)

	req, _ := http.NewRequest("GET", apiRequest, nil)

	resp, err := d.getClient().Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	type ConversionResponse struct {
		BaseCode       string  `json:"base_code"`
		TargetCode     string  `json:"target_code"`
		ConversionRate float64 `json:"conversion_rate"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var conversionResponse ConversionResponse
	jsonErr := json.Unmarshal(body, &conversionResponse)
	if jsonErr != nil {
		return 0, jsonErr
	}

	//fmt.Println(conversionResponse)

	return conversionResponse.ConversionRate, nil
}

type MockRateService struct {
}

func (m *MockRateService) GetExchangeRate(base string, to string) (float64, error) {
	return 1.5, nil
}
