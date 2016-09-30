package main

import (
	"encoding/json"
	"fmt"
	"os"

	"code.google.com/p/gcfg"
	"github.com/kylelemons/go-gypsy/yaml"
)

// yaml
func yamlConf() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}

type configuration struct {
	Enabled bool
	Path    string
}

// json
func jsonConf() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}

// ini
func iniConf() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}

func main() {
	jsonConf()
	yamlConf()
	iniConf()
}
