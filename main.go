package main

import (
	"api"
	"data"
	"logger"
)

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Program starting...")

	data.InitTables()

	api.InitApi()
}