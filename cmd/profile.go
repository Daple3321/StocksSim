package cmd

import (
	"fmt"
	"sync"

	"github.com/Daple3321/StocksSim/stock"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")

	headerStyle  = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(14)
	oddRowStyle  = cellStyle.Foreground(gray)
	evenRowStyle = cellStyle.Foreground(lightGray)
)

// Fetches current price info for all stocks in parallel
func fetchPortfolioInfos() []*stock.StockInfo {
	if !p.HasStocks() {
		return nil
	}
	infos := make([]*stock.StockInfo, len(p.Stocks))
	var wg sync.WaitGroup
	for i := range p.Stocks {
		i := i
		ticker := p.Stocks[i].Ticker
		wg.Go(func() {
			infos[i], _ = p.Fetcher.Fetch(ticker)
		})
	}
	wg.Wait()
	return infos
}

// Computes portfolio growth and current price from pre-fetched infos
func getPortfolioStatsFromInfos(infos []*stock.StockInfo) (growth float64, currentPrice float64) {
	if len(infos) != len(p.Stocks) {
		return 0, 0
	}
	var sum float64
	var oldCost float64
	for i := range p.Stocks {
		oldCost += p.Stocks[i].OriginalCost
		if infos[i] != nil {
			sum += infos[i].Price * float64(p.Stocks[i].Amount)
		}
	}
	if oldCost == 0 {
		return 0, sum
	}
	return (sum / oldCost) - 1, sum
}

func GetPortfolioTable(infos []*stock.StockInfo) *table.Table {

	if !p.HasStocks() || len(infos) != len(p.Stocks) {
		return nil
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("Stock", "Amount", "Current price", "Growth (%)")

	// Build rows
	for i := range p.Stocks {
		s := &p.Stocks[i]
		info := infos[i]
		if info == nil {
			t.Row(s.Ticker, fmt.Sprint(s.Amount), "—", "—")
			continue
		}
		currentVal := info.Price * float64(s.Amount)
		growth := (currentVal / s.OriginalCost) - 1
		currentPriceStr := "$" + fmt.Sprintf("%.2f", info.Price*float64(s.Amount))
		growthStr := fmt.Sprintf("%.2f", growth*100) + "%"
		if growth > 0 {
			growthStr = "+" + growthStr
		}
		t.Row(s.Ticker, fmt.Sprint(s.Amount), currentPriceStr, growthStr)
	}

	return t
}

var profileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Prints player's profile",
	Args:    cobra.MaximumNArgs(0),
	Aliases: []string{"p", "pf"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(okStyle.Render(fmt.Sprintf("You have $%.2f USD", p.Usd)))

		if !p.HasStocks() {
			fmt.Println(warningStyle.Render("Currently you have no stocks."))
			return
		}

		infos := fetchPortfolioInfos()
		portfolioGrowth, currentPrice := getPortfolioStatsFromInfos(infos)
		if portfolioGrowth > 0 {
			fmt.Println(okStyle.Render(fmt.Sprintf("You have $%.2f in stocks [+%.2f]", currentPrice, portfolioGrowth)))
		} else if portfolioGrowth < 0 {
			fmt.Println(okStyle.Render(fmt.Sprintf("You have $%.2f in stocks %.2f", currentPrice, portfolioGrowth)))
		}

		fmt.Println(GetPortfolioTable(infos))
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
