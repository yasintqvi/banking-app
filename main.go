package main

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/app"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
