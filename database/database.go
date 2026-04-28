package database

import (
	"context"
	"fmt"
	"time"

	"github.com/irvanrifai/go-clean-architecture/config"
	"github.com/irvanrifai/go-clean-architecture/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Provider untuk MySQL
func NewMySQLDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Mysql.User,
		cfg.Database.Mysql.Password,
		cfg.Database.Mysql.Host,
		cfg.Database.Mysql.Port,
		cfg.Database.Mysql.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		pkg.ZapLog.Fatal("Failed to connect to mysql database: " + err.Error())
	}
	return db
}

// Provider untuk MongoDB
func NewMongoDB(cfg *config.Config) *mongo.Client {
	mongoURI := fmt.Sprintf("mongodb://%s:%s",
		cfg.Database.Mongo.Host,
		cfg.Database.Mongo.Port,
	)
	
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		pkg.ZapLog.Fatal("Failed to connect to mongo database: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		pkg.ZapLog.Fatal("MongoDB ping failed: " + err.Error())
	}

	return client
}