package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB    DBConfig     `yaml:"db"`
	HTTP  ServerConfig `yaml:"http"`
	Token TokenConfig  `yaml:"token"`
	App   AppConfig    `yaml:"app"`
}

type AppConfig struct {
	AppName    string `yaml:"app_name"`
	LogLevel   string `yaml:"log_level"`
	Production bool   `yaml:"production"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type TokenConfig struct {
	SecretKey  string        `env:"TOKEN_SECRET_KEY" env-required:"true"`
	TimeToLive time.Duration `yaml:"time_to_live"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}