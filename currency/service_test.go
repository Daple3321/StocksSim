package currency_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Daple3321/StocksSim/currency"
	"github.com/joho/godotenv"
)

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}

func Test_GetExchangeRate(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	mockClient := &http.Client{Transport: &mockTransport{
		roundTrip: func(*http.Request) (*http.Response, error) {
			body := `{"base_code":"USD","target_code":"EUR","conversion_rate":0.92}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		},
	}}

	rateService := currency.NewDefaultRateService(mockClient)

	t.Run("No base currency", func(t *testing.T) {
		_, err := rateService.GetExchangeRate("", "EUR")

		if err.Error() != "no base currency provided" {
			t.Errorf("Wrong error thrown. got: %s. wanted: %s", err, "no base currency provided")
		}
	})

	t.Run("No target currency", func(t *testing.T) {
		_, err := rateService.GetExchangeRate("USD", "")

		if err.Error() != "no target currency provided" {
			t.Errorf("Wrong error thrown. got: %s. wanted: %s", err, "no target currency provided")
		}
	})

	t.Run("No CurrencyExchange_API_KEY key", func(t *testing.T) {

		os.Unsetenv("CurrencyExchange_API_KEY")

		_, err := rateService.GetExchangeRate("USD", "EUR")

		if err.Error() != "CurrencyExchange_API_KEY is missing in environment variables" {
			t.Errorf("Wrong error thrown. got: %s. wanted: %s", err, "CurrencyExchange_API_KEY is missing in environment variables")
		}

	})

	// var tests = []struct {
	// 	name        string
	// 	input       [2]string
	// 	wantedValue float64
	// 	wantedErr   error
	// }{
	// 	{"No base currency", [2]string{"", "EUR"}, 0.0, errors.New("no base currency provided")},
	// 	{"No target currency", [2]string{"USD", ""}, 0.0, errors.New("no target currency provided")},
	// 	{"No ApiKey", [2]string{"USD", "EUR"}, 0.0, errors.New("CurrencyExchange_API_KEY is missing in environment variables")},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		ans, err := rateService.GetExchangeRate(tt.input[0], tt.input[1])

	// 		if err.Error() != tt.wantedErr.Error() {
	// 			t.Errorf("Wrong error thrown. got: %s. wanted: %s", err, tt.wantedErr)
	// 			return
	// 		}

	// 		// if !errors.Is(err, tt.wantedErr) {
	// 		// 	t.Errorf("Wrong error thrown. got: %s. wanted: %s", err, tt.wantedErr)
	// 		// 	return
	// 		// }

	// 		// if tt.wantedErr && err == nil {
	// 		// 	t.Errorf("No error thrown")
	// 		// 	return
	// 		// }

	// 		if tt.wantedValue == 0.0 {
	// 			return
	// 		}

	// 		if ans != tt.wantedValue {
	// 			t.Errorf("got %f, want %f", ans, tt.wantedValue)
	// 		}
	// 	})
	// }
}
