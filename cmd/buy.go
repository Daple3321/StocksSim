package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "buy [stock ticker name] [amount]",
	Short: "Buys specified amount of stocks",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ticker, _ := cmd.Flags().GetString("ticker")
		if ticker == "" {
			fmt.Print(errorStyle.Render("No ticker name provided\n"))
			return
		}

		amount, _ := cmd.Flags().GetInt("amount")
		if amount <= 0 {
			fmt.Print(errorStyle.Render("Wrong buy amount\n"))
			return
		}

		p.BuyStock(ticker, amount)
	},
}

func init() {
	buyCmd.Flags().StringP("ticker", "t", "", "Stock ticker")
	buyCmd.Flags().IntP("amount", "n", 0, "Amount to buy")
	rootCmd.AddCommand(buyCmd)
}
