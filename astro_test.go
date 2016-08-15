package astro

import "testing"

func TestGetApp(t *testing.T) {
	app_name := getApp("examples/basic_example.toml")
	if app_name != "astro" {
		t.Error("Expected: astro, Got: ", app_name)
	}
}
