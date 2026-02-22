package player

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/Daple3321/StocksSim/stock"
	"github.com/Daple3321/StocksSim/utils"
)

const (
	PLAYER_FILE_NAME string  = "Player.json"
	STARTING_MONEY   float64 = 1000
)

type Player struct {
	Usd    float64       `json:"usd"`
	Stocks []stock.Stock `json:"stocks"`

	Fetcher StockFetcher `json:"-"`
}

func NewPlayer() *Player {

	p := Player{
		Fetcher: &DefaultStockFetcher{},
	}

	err := p.TryLoad()
	if err != nil {
		fmt.Printf("Error loading player: %s\n", err)
	}

	return &p
}

func (p *Player) TryLoad() error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	playerPath := path.Join(wd, PLAYER_FILE_NAME)
	if !utils.CheckFileExistence(playerPath) {
		fmt.Printf("No save file found. Creating new one at: %s\n", playerPath)
		p.Usd = STARTING_MONEY
		return p.Save()
	}

	data, fileErr := os.ReadFile(playerPath)
	if fileErr != nil {
		if errors.Is(fileErr, os.ErrNotExist) {
			fmt.Printf("%s path does not exist.", playerPath)
		}
	}

	jsonErr := json.Unmarshal(data, &p)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

// Executed everytime after commands are executed (in root.go)
func (p *Player) Save() error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	playerPath := path.Join(wd, PLAYER_FILE_NAME)

	file, fileErr := os.Create(playerPath)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	playerJson, jsonErr := json.Marshal(p)
	if jsonErr != nil {
		return jsonErr
	}
	file.Write(playerJson)

	//fmt.Printf("Save path: %s\n", playerPath)

	return nil
}

func (p *Player) BuyStock(ticker string, amount int) {

	stockInfo, fetchErr := p.Fetcher.Fetch(ticker)
	if fetchErr != nil {
		fmt.Printf("Error fetching stock info: %s\n", fetchErr)
	}

	finalPrice := stockInfo.Price * float64(amount)
	if finalPrice > p.Usd {
		fmt.Printf("Not enough money {$%f} to buy %d stocks for $%.2f\n", p.Usd, amount, finalPrice)
		return
	}

	idx := slices.IndexFunc(p.Stocks, func(s stock.Stock) bool {
		return s.Ticker == ticker
	})

	if idx != -1 {
		p.Stocks[idx].Amount += amount
		p.Stocks[idx].OriginalCost += finalPrice

	} else {
		newStock := stock.Stock{Ticker: ticker, Amount: amount, OriginalCost: finalPrice}
		p.Stocks = append(p.Stocks, newStock)
	}

	p.Usd -= finalPrice
	fmt.Printf("%d %s [%s] stocks BOUGHT for $%.2f\n", amount, stockInfo.Name, stockInfo.Ticker, finalPrice)
}

func (p *Player) SellStock(ticker string, amount int) {

	idx := slices.IndexFunc(p.Stocks, func(s stock.Stock) bool {
		return s.Ticker == ticker
	})
	if idx == -1 {
		fmt.Printf("No %s ticker found in portfolio\n", ticker)
		return
	} else if p.Stocks[idx].Amount < amount {
		fmt.Printf("Not enough stocks to sell\n")
		return
	}

	stockInfo, fetchErr := p.Fetcher.Fetch(ticker)
	if fetchErr != nil {
		fmt.Printf("Error fetching stock info: %s\n", fetchErr)
	}

	sellPrice := stockInfo.Price * float64(amount)
	p.Stocks[idx].Amount -= amount
	p.Stocks[idx].OriginalCost -= sellPrice

	p.Usd += sellPrice
	fmt.Printf("%d %s [%s] stocks SOLD for $%.2f\n", amount, stockInfo.Name, stockInfo.Ticker, sellPrice)
}

// func (p *Player) GetPortfolioCurrentPrice() float64 {

// 	if !p.HasStocks() {
// 		return 0
// 	}

// 	var sum float64
// 	var wg sync.WaitGroup

// 	for _, s := range p.Stocks {
// 		wg.Go(func() {
// 			stockInfo, _ := p.Fetcher.Fetch(s.Ticker)
// 			sum += stockInfo.Price * float64(s.Amount)
// 		})
// 	}
// 	wg.Wait()

// 	return sum
// }

// func (p *Player) GetPortfolioStats() (growth float64, currentPrice float64) {

// 	if !p.HasStocks() {
// 		return 0.0, 0.0
// 	}

// 	growth = 0.0
// 	currentPrice = p.GetPortfolioCurrentPrice()

// 	var oldPrice float64
// 	for _, s := range p.Stocks {
// 		oldPrice += s.OriginalCost
// 	}
// 	growth = (currentPrice / oldPrice) - 1

// 	return growth, currentPrice
// }

func (p *Player) HasStocks() bool {

	if len(p.Stocks) <= 0 || p.Stocks == nil {
		return false
	}

	return true
}
