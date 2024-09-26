package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config contains configuration parameters for this application.
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
}

// New returns the application configuration parameters based on environment.
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

	return cfg, nil
}

// getConfigPath validates the path to the YAML configuration file and returns it.
func getConfigPath(env string) (string, error) {
	var path string
	switch env {
	case "dev":
		path = "./config/config-dev.yml"
	case "stg":
		path = "./config/config-stg.yml"
	case "prd":
		path = "./config/config-prd.yml"
	default:
		return "", fmt.Errorf("error: '%s' is not a valid environment", env)
	}

	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if stat.IsDir() {
		return "", fmt.Errorf("'%s' is a directory, not a file", path)
	}

	return path, nil
}
