package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		fmt.Println(p.GetPortfolioTable())

	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
