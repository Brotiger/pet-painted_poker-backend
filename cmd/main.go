package main

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/docs/swagger"
	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	"github.com/Brotiger/per-painted_poker-backend/internal/router"
	"github.com/Brotiger/per-painted_poker-backend/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// @title Core API
// @BasePath /api
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	app := fiber.New(fiber.Config{})

	mongodbCtx, cancelMongodbCtx := context.WithTimeout(ctx, time.Duration(config.Cfg.MongoDB.TimeoutMs)*time.Microsecond)
	defer cancelMongodbCtx()

	mongodbClient, err := mongodb.Connect(
		mongodbCtx,
		config.Cfg.MongoDB.Uri,
		config.Cfg.MongoDB.Username,
		config.Cfg.MongoDB.Password,
	)
	if err != nil {
		log.Fatalf("failed to connect to mongodb, error: %v", err)
	}

	connection.DB = mongodbClient.Database(config.Cfg.MongoDB.Database)

	router.SetupRouter(app)

	log.Info("application started")
	log.Infof("local API docs: http://%s", swagger.SwaggerInfo.Host+"/swagger/index.html")

	if err := app.Listen(config.Cfg.Fiber.Listen); err != nil {
		log.Errorf("failed to listen, error: %v", err)
	}
}
