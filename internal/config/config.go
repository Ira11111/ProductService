package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env    string       `yaml:"env"`
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"http"`
}

type DBConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string
	Name string `yaml:"database"`
	SSL  string `yaml:"ssl_mode"`
}

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	Handler      string        `yaml:"handler"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func MustLoad() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("CONFIG_PATH environment variable not set")
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		panic("DB_PASS environment variable not set")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	cfg.DB.Pass = dbPass
	return &cfg, nil
}
