package main

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
	fmt.Printf("Application => %v\n", app)
}
