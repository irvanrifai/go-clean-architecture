package main

import (
	"github.com/irvanrifai/go-clean-architecture/database"
	"github.com/irvanrifai/go-clean-architecture/pkg"
)

func main() {
	pkg.InitLog()
	defer pkg.ZapLog.Sync()
	
	database.ConnectDB()
	// db := database.DB

	pkg.ZapLog.Info("Connected")
}
