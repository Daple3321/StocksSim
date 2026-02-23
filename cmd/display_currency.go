package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var displayCurrencyCmd = &cobra.Command{
	Use:   "currency [code]",
	Short: "Changes display currency. (accepts 3 letter currency codes like: USD, EUR, RUB)",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		newCode, _ := cmd.Flags().GetString("code")
		if newCode == "" {
			fmt.Println(errorStyle.Render("No currency code provided\n"))
			return
		}

		p.DisplayCurrency = newCode

		fmt.Println(okStyle.Render("Display currency changed to", newCode))
	},
}

func init() {
	displayCurrencyCmd.Flags().StringP("code", "c", "", "Currency code (USD, EUR, RUB, ...)")
	rootCmd.AddCommand(displayCurrencyCmd)
}
