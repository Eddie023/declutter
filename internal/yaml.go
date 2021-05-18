package internal

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// This makes output key-value in yaml generic
type Conf struct {
	Output map[string][]string
}

func (c *Conf) GetConf() *Conf {
	configFilePATH, _ := filepath.Abs("./config.yaml")

	yamlFile, err := ioutil.ReadFile(configFilePATH)
	if err != nil {
		log.Printf("Error: YamlFile.Get arr #%v ", err)

		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Error: Unmarshal Failed: %v", err)

		os.Exit(1)
	}

	return c
}
