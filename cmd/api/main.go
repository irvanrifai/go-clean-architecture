package main

import (
	"github.com/irvanrifai/go-clean-architecture/config"
	"github.com/irvanrifai/go-clean-architecture/database"
	"github.com/irvanrifai/go-clean-architecture/pkg"
)

func main() {
	config.InitConfig()
	pkg.InitLog()
	defer pkg.ZapLog.Sync()
	
	database.ConnectDB()
	// db := database.DB

	pkg.ZapLog.Info("Connected")
}
