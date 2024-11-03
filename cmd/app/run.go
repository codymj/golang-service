package app

import (
	"flag"

	"github.com/rs/zerolog/log"
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
		log.Error().Msg(err.Error())
		return
	}

	// Start application.
	log.Info().Msgf("application environment: %s", env)
	app := application{
		cfg: cfg,
	}
	app.start()
}
