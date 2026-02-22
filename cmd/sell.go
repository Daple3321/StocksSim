package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sellCmd = &cobra.Command{
	Use:   "sell [stock ticker name] [amount]",
	Short: "Sells specified amount of stocks",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ticker, _ := cmd.Flags().GetString("ticker")
		if ticker == "" {
			fmt.Print(errorStyle.Render("No ticker name provided\n"))
			return
		}

		amount, _ := cmd.Flags().GetInt("amount")
		if amount <= 0 {
			fmt.Print(errorStyle.Render("Wrong sell amount\n"))
			return
		}

		p.SellStock(ticker, amount)
	},
}

func init() {
	sellCmd.Flags().StringP("ticker", "t", "", "Stock ticker")
	sellCmd.Flags().IntP("amount", "n", 0, "Amount to sell")
	rootCmd.AddCommand(sellCmd)
}
