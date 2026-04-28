package database

import (
	"fmt"

	"github.com/irvanrifai/go-clean-architecture/config"
	"github.com/irvanrifai/go-clean-architecture/pkg"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		pkg.ZapLog.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().Database.Mysql.User,
		config.GetConfig().Database.Mysql.Password,
		config.GetConfig().Database.Mysql.Host,
		config.GetConfig().Database.Mysql.Port,
		config.GetConfig().Database.Mysql.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		pkg.ZapLog.Fatal("Failed to connect to database: " + err.Error())
	}

	pkg.ZapLog.Info("Database connection successful")
	DB = db
}
