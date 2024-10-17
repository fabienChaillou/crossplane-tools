package core

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const FORMAT = "FORMAT"

var DataConfig Config

func init() {
	loadConfigs()
	// is_format := os.Getenv(FORMAT)
	// if is_format
}

func loadConfigs() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := read(dir + "/cmd/config/config.yaml")

	if err != nil {
		log.Fatal("Error while reading app config file", err)
	}

	laodServerConfig(data)

}

func laodServerConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &DataConfig)
	fmt.Println("Loaded Server Config with ", len(DataConfig.Config), " config files")
}
