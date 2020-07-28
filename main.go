package main

import (
	"github.com/ahmadrezamusthafa/assessment/common/container"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/pkg/cache"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	"github.com/ahmadrezamusthafa/assessment/server"
	httphealth "github.com/ahmadrezamusthafa/assessment/server/http/health"
)

func main() {
	logger.SetupLogger()
	conf, err := config.New()
	if err != nil {
		logger.Warn("%v", err)
	}

	logger.Info("Starting service container...")
	container := container.NewContainer()
	container.RegisterService("config", *conf)
	container.RegisterService("database", new(database.AssessmentDatabase))
	container.RegisterService("cache", new(cache.AssessmentCache))
	container.RegisterService("magazineService", new(magazinegun.MagazineService))
	container.RegisterService("healthHandler", new(httphealth.Handler))

	rootHandler := new(server.RootHandler)
	container.RegisterService("rootHandler", rootHandler)
	if err := container.Ready(); err != nil {
		logger.Fatal("Failed to populate services %v", err)
	}

	httpServer := server.HttpServer{Config: *conf, RootHandler: rootHandler}
	httpServer.Serve()
}
