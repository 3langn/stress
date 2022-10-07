package main

import (
	"variant-service/config"
	"variant-service/internal/controller"
	"variant-service/internal/http"
	"variant-service/internal/lib/db"
	"variant-service/internal/lib/logger"
	"variant-service/internal/repository"
	"variant-service/internal/service"
	"variant-service/utils"

	"go.uber.org/fx"
)

// @title variant-service restful API
// @version 1.0.0
// @description variant-service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://localhost:8080
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func inject() fx.Option {
	return fx.Options(
		//fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		//	return &fxevent.ZapLogger{Logger: logger}
		//}),
		fx.NopLogger,
		fx.Provide(
			config.NewConfig,
			utils.NewTimeoutContext,
		),
		db.GormModule,
		logger.LoggerModule,
		repository.Module,
		service.Module,
		controller.Module,
		//nsq.ProducerModule,
		//nsq.ConsumerModule,
		http.Module,
	)
}

func main() {
	fx.New(inject()).Run()
}
