package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Host       string `toml:"host" toml-required:"true"`
	Port       string `toml:"port" toml-required:"true"`
	TryPattern string `toml:"try_pattern" toml-required:"true"`
}

type Websites []*Config

func New() *Config {
	return &Config{}
}

func (c *Config) Read(path string) error {
	err := cleanenv.ReadConfig(path, c)
	if err != nil {
		return err
	}

	return nil
}
