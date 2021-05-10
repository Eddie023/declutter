package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// FIXME: Fix absolute path problem
var configFilePATH, _ = filepath.Abs("internal/../config.yaml")

// This makes output key-value in yaml generic
type Conf struct {
	Output map[string][]string
}

func (c *Conf) GetConf() *Conf {
	fmt.Println(configFilePATH)
	yamlFile, err := ioutil.ReadFile(configFilePATH)
	if err != nil {
		log.Printf("yamlFile.Get arr #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
