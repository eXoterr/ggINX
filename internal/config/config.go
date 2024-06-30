package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP
	Logger
}

type HTTP struct {
	ReadTimeout    uint `toml:"read_timeout" env-default:"5"`
	WriteTimeout   uint `toml:"write_timeout" env-default:"5"`
	MaxConnections uint `toml:"max_connections" env-default:"30"`
}

type Logger struct {
	OutType       string `toml:"log_out_type" toml-required:"true"`
	LogLocation   string `toml:"log_location" toml-required:"true"`
	Level         string `toml:"log_level" toml-required:"true"`
	WithTimestamp bool   `toml:"log_timestamp" toml-required:"true"`
}

func (conf *Config) Setup(confiPath string) error {
	err := cleanenv.ReadEnv(conf)
	if err != nil {
		return err
	}

	err = cleanenv.ReadConfig(confiPath, conf)
	if err != nil {
		return err
	}

	return nil
}

func New() *Config {
	return &Config{}
}
