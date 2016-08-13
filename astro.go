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
	version := config.Get("application.version")
	get := config.Get("build.go.get")
	//build_go_get := config.Get("build.go")
	fmt.Printf("Application => %v %v\n", app, version)
	fmt.Println("Build => ", get)
}
