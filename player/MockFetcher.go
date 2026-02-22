package player

import (
	"time"

	"github.com/Daple3321/StocksSim/stock"
	"github.com/Daple3321/StocksSim/utils"
)

type MockFetcher struct{}

func (m *MockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

	stockChannel := make(chan *stock.StockInfo)

	go func() {
		time.Sleep(450 * time.Millisecond)

		stockChannel <- &stock.StockInfo{
			Ticker:   ticker,
			Name:     "MockName",
			Price:    utils.RandFloat(1, 250),
			Exchange: "NASDAQ",
			Updated:  time.Now().Unix(),
			Currency: "USD",
		}
	}()

	return <-stockChannel, nil
}
