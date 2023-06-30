package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./yaml2toml <inputfile>")
		os.Exit(1)
	}

	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		panic(err)
	}

	tomlBytes, err := toml.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(tomlBytes))
}
