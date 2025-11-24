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

	// 1. Загружаем переменные из .env файла
	// Если файла нет (например, на сервере), программа продолжит работу
	if err := godotenv.Load(); err != nil {
		// Это не фатальная ошибка, если мы запускаем не локально,
		// но для отладки полезно знать.
		log.Println("No .env file found")
	}

	cmd.Execute()
}
