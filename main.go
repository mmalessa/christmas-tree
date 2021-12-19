// Copyright 2021 Marcin Malessa

package main

import (
	"fmt"

	tree "christmastree/pkg/christmastree"

	"github.com/spf13/viper"
)

var (
	cfgDir string = "/etc/christmastree"
)

func main() {
	fmt.Printf("Loading config from: %s\n", cfgDir)
	LoadConfig(cfgDir)

	gpioPin := viper.GetInt("strip.gpiopin")
	ledCount := viper.GetInt("strip.ledcount")
	brightness := viper.GetInt("strip.brightness")
	fmt.Printf("Start with gpioPin %d, ledCount %d, brightness %d\n", gpioPin, ledCount, brightness)

	tr := tree.NewChristmasTree(gpioPin, ledCount, brightness)
	defer tr.Defer()

	patterns := viper.GetStringMap("patterns")
	for patternid, patternconfig := range patterns {
		configmap := patternconfig.(map[string]interface{})
		templatename := configmap["template"].(string)
		config := configmap["config"].(map[string]interface{})
		err := tr.AddPattern(patternid, templatename, config)
		if err != nil {
			panic(err)
		}
	}

	yamlmatrix := viper.Get("treematrix")
	buildTreeMatrix(tr, yamlmatrix)

	playlist := viper.Get("playlist").([]interface{})
	for {
		for _, pattern := range playlist {
			fmt.Printf("Play pattern id: %s\n", pattern.(string))
			tr.PlayPattern(pattern.(string))
		}
	}
}

func buildTreeMatrix(tr *tree.ChristmasTree, yamlmatrix interface{}) error {
	var matrix map[int]map[int]int
	matrix = make(map[int]map[int]int)
	rows := yamlmatrix.([]interface{})
	for nr, row := range rows {
		matrix[nr] = make(map[int]int)
		columns := row.([]interface{})
		for nc, value := range columns {
			matrix[nr][nc] = value.(int)
		}
	}
	tr.SetMatrix(matrix)
	return nil
}
