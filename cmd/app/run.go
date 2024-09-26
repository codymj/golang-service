package app

import (
	"flag"
	"log/slog"
	"os"

	"golang-service.codymj.io/config"
)

// Run gathers configuration data then runs the application.
func Run() {
	// Parse commandline flags.
	var env string
	flag.StringVar(&env, "env", "dev", "environment: dev|stg|prd")
	flag.Parse()

	// Parse environment-based configuration file.
	cfg, err := config.New(env)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// Start application.
	app := application{
		cfg:    cfg,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	app.start()
}
