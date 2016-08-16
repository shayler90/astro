package application

import "testing"

func TestGetName(t *testing.T) {
	app_name := getName("examples/basic_example.toml")
	if app_name != "astro" {
		t.Error("Expected: astro, Got: ", app_name)
	}
}

func TestGetSource(t *testing.T) {
	app_source := getSource("examples/basic_example.toml")
	if app_source != "https://github.com/dnt17/astro" {
		t.Error("Expected: https://github.com/dnt17/astro, Got: ", app_source)
	}
}

func TestGetPullMaster(t *testing.T) {
	app_pull_master := getPullMaster("examples/basic_example.toml")
	if app_pull_master != true {
		t.Error("Expected: true, Got: ", app_pull_master)
	}
}

func TestGetVersion(t *testing.T) {
	app_version := getVersion("examples/basic_example.toml")
	if app_version != "0.0.1" {
		t.Error("Expected: 0.0.1, Got: ", app_version)
	}
}
