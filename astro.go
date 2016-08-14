package astro

import (
	"github.com/pelletier/go-toml"
	"log"
)

func getApp() string {
	config, err := toml.LoadFile("examples/basic_example.toml")
	if err != nil {
		log.Fatal(err)
	}
	var app_name string
	app_name = config.Get("application.name").(string)
	return app_name
}
