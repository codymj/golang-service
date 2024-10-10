package app

import (
	"flag"
	"log/slog"

	"golang-service.codymj.io/configs"
)

// Gathers configuration data then runs the application.
func Run() {
	// Parse commandline flags.
	var env string
	flag.StringVar(&env, "env", "dev", "environment: dev|stg|prd")
	flag.Parse()

	// Parse environment-based configuration file.
	cfg, err := configs.New(env)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// Start application.
	app := application{
		cfg: cfg,
	}
	app.start()
}
