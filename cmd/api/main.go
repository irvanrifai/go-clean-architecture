package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/irvanrifai/go-clean-architecture/config"
	"github.com/irvanrifai/go-clean-architecture/database"
	"github.com/irvanrifai/go-clean-architecture/internal/delivery/http"
	"github.com/irvanrifai/go-clean-architecture/internal/delivery/http/handler"
	"github.com/irvanrifai/go-clean-architecture/internal/repository"
	"github.com/irvanrifai/go-clean-architecture/internal/usecase"
	"github.com/irvanrifai/go-clean-architecture/pkg"
	"go.uber.org/fx"
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
			repository.NewUserRepository,
			usecase.NewUserUsecase,
			handler.NewUserHandler,
			http.NewRouter,
        ),
        fx.Invoke(func(lc fx.Lifecycle, r *gin.Engine) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go r.Run(":8080")
					return nil
				},
			})
		}),
    ).Run()
}
