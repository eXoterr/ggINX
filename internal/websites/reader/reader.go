package reader

import (
	"os"
	"path/filepath"

	"github.com/eXoterr/ggINX/internal/websites/config"
)

type ConfigReader struct{}

func New() *ConfigReader {
	return &ConfigReader{}
}

func (r *ConfigReader) Read(path string) (config.Websites, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	websites := config.Websites{}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".toml" {
			cfg := config.New()
			err := cfg.Read(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}
			websites = append(websites, cfg)
		}
	}

	return websites, nil
}
