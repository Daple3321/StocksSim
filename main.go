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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cmd.Execute()
}
