package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const filePath = "src/config/urlConfig.yaml"

func ConfigLoader() map[interface{}]map[interface{}]string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var config = make(map[interface{}]map[interface{}]string)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
