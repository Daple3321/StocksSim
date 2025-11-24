package cmd

import (
	"fmt"

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

func GetPortfolioTable() *table.Table {

	if !p.HasStocks() {
		return nil
	}

	//sb := strings.Builder{}

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

	for _, s := range p.Stocks {

		growth := s.GetStockGrowth()
		if growth > 0 {
			t.Row(s.Ticker, fmt.Sprint(s.Amount), "$"+fmt.Sprintf("%.2f", s.OriginalCost+(s.OriginalCost*growth)), "+"+fmt.Sprintf("%.2f", growth*100)+"%")

		} else if growth <= 0 {

			t.Row(s.Ticker, fmt.Sprint(s.Amount), "$"+fmt.Sprintf("%.2f", s.OriginalCost+(s.OriginalCost*growth)), fmt.Sprintf("%.2f", growth*100)+"%")
		}

	}

	return t
}

var profileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Prints player's profile",
	Args:    cobra.MaximumNArgs(0),
	Aliases: []string{"p", "pf"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf(okStyle.Render("You have $%.2f USD\n"), p.Usd)
		portfolioGrowth := p.GetPortfolioGrowth()
		if portfolioGrowth > 0 {
			fmt.Printf(okStyle.Render("You have $%.2f in stocks [+%.2f]\n"), p.GetPortfolioCurrentPrice(), portfolioGrowth)
		} else {
			fmt.Printf(okStyle.Render("You have $%.2f in stocks [%.2f]\n"), p.GetPortfolioCurrentPrice(), portfolioGrowth)
		}
		fmt.Print(("Portfolio\n"))
		fmt.Println(GetPortfolioTable())

	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
