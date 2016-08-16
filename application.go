package application

import (
	"github.com/pelletier/go-toml"
	"log"
)

type Application struct {
	name        string
	source      string
	pull_master bool
	version     string
}

func getName(tomlfile string) string {
	tree, err := toml.LoadFile(tomlfile)
	if err != nil {
		log.Fatal(err)
	}
	var app_name string
	app_name = tree.Get("application.name").(string)
	return app_name
}

func getSource(tomlfile string) string {
	tree, err := toml.LoadFile(tomlfile)
	if err != nil {
		log.Fatal(err)
	}
	var app_source = tree.Get("application.source").(string)
	return app_source
}
