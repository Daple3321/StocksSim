/*
Copyright © 2025 Daple <GameRoll>
*/
package main

import (
	"log"

	"gameroll.com/StocksSim/cmd"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cmd.Execute()
}
