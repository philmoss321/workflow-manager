package main

import (
	"log"

	"github.com/BurntSushi/toml"
	model "github.com/philmoss321/workflow-manager/core/model"
	factory "github.com/philmoss321/workflow-manager/factory"
)

// Version : App version
var Version string

func main() {
	config, err := setConfig()
	log.Println(config)
	if err != nil {
		log.Fatal(err.Error() + ". Can't set configuration values, Packager shutting down")
		return
	}
	// Run(config)
	workflow := factory.CreatePremadeWorkflow("PackageHLS")
	err = workflow.Run()
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Fatal("Error initializing, Packager shutting down")
		return
	}
}

func setConfig() (*model.Config, error) {
	var conf model.Config
	if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
