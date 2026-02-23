package cmd

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	checkBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("99")).
			Padding(1, 2).
			MarginTop(1)
	checkLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Width(10)
	checkValueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("188"))
	checkTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true)
)

var checkCmd = &cobra.Command{
	Use:   "check [stock ticker name]",
	Short: "Fetches ticker data",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var content string

		if len(args) == 0 {
			content = errorStyle.Render("No ticker name provided")
			fmt.Println(checkBoxStyle.Render(content))
			return
		}

		stockInfo, err := p.Fetcher.Fetch(args[0])
		if err != nil {
			content = errorStyle.Render(fmt.Sprintf("Error fetching stock info: %s", err))
			fmt.Println(checkBoxStyle.Render(content))
			return
		}

		// check if needs conversion
		// convert to display currency
		if p.DisplayCurrency != BaseCurrency {
			usdPrice := stockInfo.Price
			stockInfo.ConvertedPrice = p.Converter.Convert("USD", p.DisplayCurrency, usdPrice)
		}

		line := func(label, value string) string {
			return checkLabelStyle.Render(label+":") + " " + checkValueStyle.Render(value)
		}
		t := time.Unix(stockInfo.Updated, 0)
		title := checkTitleStyle.Render(fmt.Sprintf("%s [%s]", stockInfo.Name, stockInfo.Ticker))
		content = title + "\n\n" +
			line("Price", fmt.Sprintf("%.2f", stockInfo.ConvertedPrice)) + "\n" +
			line("Exchange", stockInfo.Exchange) + "\n" +
			line("Updated", t.String()) + "\n" +
			line("Currency", stockInfo.Currency)

		fmt.Println(checkBoxStyle.Render(content))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
