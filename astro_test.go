package astro

import "testing"

func TestGetApp(t *testing.T) {
	config, err := toml.LoadFile("examples/basic_example.toml")
	if err != nil {
		t.Error(err)
	}
	app := getApp(config)
	if app != "astro" {
		t.Error("Expected 'astro', Got '%v'", app)
	}
}
