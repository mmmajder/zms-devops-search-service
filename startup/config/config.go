package config

import "os"

type Config struct {
	Port             string
	SearchDBHost     string
	SearchDBPort     string
	SearchDBUsername string
	SearchDBPassword string
	BookingHost      string
	BookingPort      string
}

func NewConfig() *Config {
	return &Config{
		Port:             os.Getenv("SERVICE_PORT"),
		SearchDBHost:     os.Getenv("DB_HOST"),
		SearchDBPort:     os.Getenv("DB_PORT"),
		SearchDBUsername: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		SearchDBPassword: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		BookingHost:      os.Getenv("BOOKING_HOST"),
		BookingPort:      os.Getenv("BOOKING_PORT"),
	}
}
