package player

import (
	"time"

	"gameroll.com/StocksSim/stock"
	"gameroll.com/StocksSim/utils"
)

type MockFetcher struct{}

func (m *MockFetcher) Fetch(ticker string) (*stock.StockInfo, error) {

	stockChannel := make(chan *stock.StockInfo)

	go func() {
		time.Sleep(450 * time.Millisecond)

		stockChannel <- &stock.StockInfo{
			Ticker: ticker,
			Name:   "MockName",
			Price:  utils.RandFloat(1, 250),
		}
	}()

	return <-stockChannel, nil
}
