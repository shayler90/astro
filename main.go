package astro

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
)

func main() {
	config, err := toml.LoadFile("examples/basic_example.toml")
	if err != nil {
		log.Fatal(err)
	}

	app := config.Get("application.name")
	version := config.Get("application.version")
	fmt.Println("Application => ", app, version)
}
