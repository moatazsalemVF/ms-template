package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	Server struct {
		Port    int
		Address string
	}
	App struct {
		Name string
	}
}

var Conf Configs

func ReadConfigs() {
	yfile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &Conf)
	if err2 != nil {

		log.Fatal(err2)
	}

	if Debug {
		d, err := yaml.Marshal(&Conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- Configs:\n%s\n\n", string(d))
	}
}
