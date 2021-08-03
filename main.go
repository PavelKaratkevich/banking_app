package main

import (
	"banking/app"
	"banking/logger"
)

func main() {

	logger.Info("Server is running...")
	app.Start()
}
