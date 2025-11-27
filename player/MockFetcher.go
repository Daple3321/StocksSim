package player

import (
	"time"

	"gameroll.com/StocksSim/stock"
)

type MockFetcher struct{}

func (m *MockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

	stockChannel := make(chan *stock.StockInfo)

	go func() {
		time.Sleep(8000)

		stockChannel <- &stock.StockInfo{
			Ticker: "AAPL",
			Name:   "Apple Inc",
			Price:  150.0,
		}
	}()

	return <-stockChannel, nil
}
