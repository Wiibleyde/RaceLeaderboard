package main

import "logger"

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Program starting...")
}