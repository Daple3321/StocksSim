package cmd

import (
	"fmt"

	"time"

	"gameroll.com/StocksSim/stock"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [stock ticker name]",
	Short: "Fetches ticker data",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Print(errorStyle.Render("No ticker name provided\n"))
			return
		}

		stockInfo, err := stock.FetchStockInfo(args[0])
		if err != nil {
			fmt.Printf(errorStyle.Render("Error fetching stock info: %s\n"), err)
		}

		fmt.Printf("%s [%s]\n", stockInfo.Name, stockInfo.Ticker)
		fmt.Printf("Price: %.2f\n", stockInfo.Price)
		fmt.Printf("Exchange: %s\n", stockInfo.Exchange)

		t := time.Unix(stockInfo.Updated, 0)
		//fmt.Println("Time from Unix seconds:", t)
		fmt.Printf("Updated: %s\n", t)

		fmt.Printf("Currency: %s\n", stockInfo.Currency)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
