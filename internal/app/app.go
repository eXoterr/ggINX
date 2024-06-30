package app

import (
	"os"

	"github.com/eXoterr/ggINX/internal/config"
	"github.com/eXoterr/ggINX/pkg/logger"
)

type Application struct {
	config *config.Config
	logger logger.Logger
}

func New() *Application {
	conf := config.New()
	logger := logger.New()

	return &Application{
		config: conf,
		logger: logger,
	}
}

func (app *Application) Start(configPath string) error {
	err := app.config.Setup(configPath)
	if err != nil {
		return err
	}

	logWriter, err := os.OpenFile(app.config.Logger.LogLocation, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	loggerCfg := &logger.Config{
		OutType:       app.config.Logger.OutType,
		Level:         app.config.Logger.Level,
		WithTimestamp: app.config.Logger.WithTimestamp,
		Writer:        logWriter,
	}

	err = app.logger.Setup(loggerCfg)
	if err != nil {
		return err
	}

	app.logger.Info("app started")

	return nil
}
