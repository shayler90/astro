package astro

import (
	"github.com/pelletier/go-toml"
	"log"
)

func getApp(tomlfile string) string {
	tree, err := toml.LoadFile(tomlfile)
	if err != nil {
		log.Fatal(err)
	}
	var app_name string
	app_name = tree.Get("application.name").(string)
	return app_name
}
