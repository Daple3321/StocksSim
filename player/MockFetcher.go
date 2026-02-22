package player

import (
	"time"

	"gameroll.com/StocksSim/stock"
)

type MockFetcher struct{}

func (m *MockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

	stockChannel := make(chan *stock.StockInfo)

	go func() {
		time.Sleep(450 * time.Millisecond)

		stockChannel <- &stock.StockInfo{
			Ticker: ticker,
			Name:   "MockName",
			Price:  50.0,
		}
	}()

	return <-stockChannel, nil
}
