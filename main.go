package main

import (
	"github.com/mvpoyatt/go-api/api"
	"github.com/mvpoyatt/go-api/configs"
	"github.com/mvpoyatt/go-api/database"
	"github.com/mvpoyatt/go-api/utils/logger"
)

func main() {
	// Read from config files
	configs.LoadConfigs()

	// Set logging level
	logger.SetLevel(configs.Values.LoggerLevel)

	// Open DB connection
	database.Connect(configs.Values.Database)

	// Start server
	api.StartServer(configs.Values.Server)
}
