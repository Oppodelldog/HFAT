package config

import (
	"HFAT/server"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type Config struct{
	ForwardingTargets []server.ForwardingTarget
	Port int
}

func ReadConfigFromFile() Config {

	file, fileErr := ioutil.ReadFile("./config.json")
	if fileErr != nil {
		fmt.Printf("File error: %v\n", fileErr)
		os.Exit(1)
	}

	var jsonType Config
	jsonErr := json.Unmarshal(file, &jsonType)
	if(jsonErr != nil){
		fmt.Printf("Json parsing error: %v\n", jsonErr)
		os.Exit(1)
	}

	return jsonType
}
