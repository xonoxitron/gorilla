package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var version = "0.1.0"

type Config struct {
	Currency string
	Token    string
	Interval int
	Debug    bool
}

func Get() *Config {
	if godotenv.Load(".env") != nil {
		panic("Error loading .env file")
	}

	currency := os.Getenv("CURRENCY")
	token := os.Getenv("TOKEN")
	interval, _ := strconv.Atoi(os.Getenv("INTERVAL"))
	debug, _ := strconv.ParseBool(strings.ToLower(os.Getenv("DEBUG")))

	return &Config{currency, token, interval, debug}
}

func Version() string {
	return version
}
