package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/viper"
)

func LoadConfig(configPath string) {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	readConfigsFromDirectory(configPath)
	cfgSubDir := filepath.Join(configPath, "conf.d")
	if _, err := os.Stat(cfgSubDir); !os.IsNotExist(err) {
		readConfigsFromDirectory(cfgSubDir)
	}
}

func readConfigsFromDirectory(directory string) {
	viper.AddConfigPath(directory)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileName := file.Name()
		if match, _ := regexp.MatchString(`[0-9]{2}.+\.yaml$`, fileName); match {
			fmt.Printf("Use config file: %s/%s\n", directory, fileName)
			viper.SetConfigName(fileName)
			viper.MergeInConfig()
		}
	}
}
