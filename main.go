/*
Copyright © 2025 Daple <GameRoll>
*/
package main

import (
	"log"

	"github.com/Daple3321/StocksSim/cmd"
	"github.com/joho/godotenv"
)

func main() {

	// TODO: IMPORTANT!! Remove these env variables.
	// Make a config system that will save to persistent user directory
	// Also need to make some kinda setup pipeline on the first launch. (When no cfg detected)
	// To assign API keys and other stuff
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cmd.Execute()
}
