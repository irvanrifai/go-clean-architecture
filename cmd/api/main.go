package main

import (
	"github.com/irvanrifai/go-clean-architecture/config"
	"github.com/irvanrifai/go-clean-architecture/database"
	"github.com/irvanrifai/go-clean-architecture/pkg"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	config.InitConfig()
	pkg.InitLog()
	defer pkg.ZapLog.Sync()
	
	fx.New(
        fx.Provide(
            config.GetConfig,
            database.NewMySQLDB,
            database.NewMongoDB,
            // ... provider lain (repository, service)
        ),
        fx.Invoke(func(db *gorm.DB, m *mongo.Client) {
            pkg.ZapLog.Info("All databases are ready!")
        }),
    ).Run()
}
