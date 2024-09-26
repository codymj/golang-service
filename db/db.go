package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang-service.codymj.io/config"
)

func New(cfg *config.Config) (*sql.DB, error) {
	// Check for .env file. If one isn't found, continue.
	_ = godotenv.Load(".env")

	// Parse environment variables.
	host, found := os.LookupEnv("MARIADB_HOST")
	if !found {
		return nil, errors.New("MARIADB_HOST environment variable is not set")
	}
	port, found := os.LookupEnv("MARIADB_PORT")
	if !found {
		return nil, errors.New("MARIADB_PORT environment variable is not set")
	}
	user, found := os.LookupEnv("MARIADB_USER")
	if !found {
		return nil, errors.New("MARIADB_USER environment variable is not set")
	}
	password, found := os.LookupEnv("MARIADB_PASSWORD")
	if !found {
		return nil, errors.New("MARIADB_PASSWORD environment variable is not set")
	}
	database, found := os.LookupEnv("MARIADB_DATABASE")
	if !found {
		return nil, errors.New("MARIADB_DATABASE environment variable is not set")
	}

	// Build DSN.
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		database,
	)

	// Open a connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.Database.MaxIdleTime)

	// Verify connection.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
