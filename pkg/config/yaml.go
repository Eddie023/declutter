package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// This makes output key-value in yaml generic
type Conf struct {
	Output map[string][]string
}

// Get the mapping of output folders and it's file types.
func (c *Conf) GetConf() *Conf {
	configFilePATH, _ := filepath.Abs("./config.yaml")

	yamlFile, err := os.ReadFile(configFilePATH)
	if err != nil {
		log.Fatalf("Error: YamlFile.Get arr #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Error: Unmarshal Failed: %v", err)
	}

	return c
}
