package main

import (
	"flag"

	"github.com/eXoterr/ggINX/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "config/config.toml", "Global server configuration path")
	flag.Parse()
}

func main() {
	app := app.New()
	err := app.Start(configPath)
	if err != nil {
		panic(err)
	}
}
