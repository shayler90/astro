package application

import "testing"

func TestGetName(t *testing.T) {
	app_name := getName("examples/basic_example.toml")
	if app_name != "astro" {
		t.Error("Expected: astro, Got: ", app_name)
	}
}
