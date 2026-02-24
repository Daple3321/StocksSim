/*
Copyright © 2025 Daple <GameRoll>
*/
package cmd

import (
	"os"

	"github.com/Daple3321/StocksSim/player"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var mainStyle = lipgloss.NewStyle()

// var defaultStyle = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("188")).
// 	Inherit(mainStyle)

var okStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("120")).
	AlignHorizontal(lipgloss.Left)
	// Inherit(mainStyle)
	// Bold(true).

var warningStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("21")).
	AlignHorizontal(lipgloss.Left)

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("124")).
	Inherit(mainStyle)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "StocksSim",
	Short: "Simulates investing in stocks & crypto",
	Long:  `...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var p player.Player

const BaseCurrency string = "USD"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	p = *player.NewPlayer()

	//fmt.Println("Test conversion:", p.Converter.Convert("USD", "RUB", 5))

	err := rootCmd.Execute()
	p.Save()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.StocksSim.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
