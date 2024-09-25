package app

import (
	"flag"
	"log/slog"
	"os"

	"golang-service.codymj.io/config"
)

func Run() {
	// Parse commandline flags.
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "./config/config.yml", "path to config.yml file")
	flag.Parse()

	// Validate configuration file path.
	if err := config.ValidateConfigPath(cfgPath); err != nil {
		slog.Error(err.Error())
	}

	// Parse configuration from file path.
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		slog.Error(err.Error())
	}

	// Start application.
	app := application{
		cfg:    cfg,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	app.serve()
}
