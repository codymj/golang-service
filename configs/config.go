package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// Contains configuration parameters for this application.
type Config struct {
	App struct {
		Namespace string `yaml:"namespace"`
		Name      string `yaml:"name"`
		Version   string `yaml:"version"`
	} `yaml:"app"`
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Read   time.Duration `yaml:"read"`
			Write  time.Duration `yaml:"write"`
			Idle   time.Duration `yaml:"idle"`
			Server time.Duration `yaml:"server"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		MaxOpenConns int           `yaml:"max_open_conns"`
		MaxIdleConns int           `yaml:"max_idle_conns"`
		MaxIdleTime  time.Duration `yaml:"max_idle_time"`
	} `yaml:"database"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

// Returns the application configuration parameters based on environment.
func New(env string) (*Config, error) {
	// Validate and get configuration file path.
	path, err := getConfigPath(env)
	if err != nil {
		return nil, err
	}

	// Open configuration file.
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode YAML file.
	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	// Initialize logger.
	initLogger(cfg)

	return cfg, nil
}

// Validates path to the YAML configuration file and returns it.
func getConfigPath(env string) (string, error) {
	// Set expected path based on environment.
	var path string
	switch env {
	case "dev":
		path = "./configs/config-dev.yml"
	case "stg":
		path = "./configs/config-stg.yml"
	case "prd":
		path = "./configs/config-prd.yml"
	default:
		return "", fmt.Errorf("error: '%s' is not a valid environment", env)
	}

	// Verify the path to the file exists.
	stat, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("invalid configuration path: %v", err)
	}
	if stat.IsDir() {
		return "", fmt.Errorf("'%s' is a directory, not a file", path)
	}

	return path, nil
}

// Initializes logger.
func initLogger(cfg *Config) {
	switch cfg.Log.Level {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		log.Warn().Msg("invalid log level in config, defaulting to info")
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
}
