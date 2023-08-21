package scrapers

import (
	"log"
	"shampoo-scraper/src/utils"
)

const readmConfigKey = "readmConfig"

var readmConfig interface{}

func loadReadmConfig() {
	readmConfig = utils.ConfigLoader()[readmConfigKey]
	log.Println("Config File Loaded!!\n")
}

func checkReadmConfigFileIsLoaded() {
	if readmConfig == nil {
		loadReadmConfig()
	}
	return
}
