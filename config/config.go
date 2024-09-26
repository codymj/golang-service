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

// New returns the application configuration parameters.
func New(configPath string) (*Config, error) {
	// Open configuration file.
	file, err := os.Open(configPath)
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

// ValidateConfigPath validates the path to the YAML configuration file.
func ValidateConfigPath(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", path)
	}

	return nil
}
