package config

import (
	"HFAT/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct represents HFAT configuration
// and is used for (de-) serialization from and to config file
type Config struct {
	ForwardingTargets []server.ForwardingTarget
	Port              int
}

// ReadConfigFromFile loads the HFAT config from disk
func ReadConfigFromFile() Config {

	file, fileErr := ioutil.ReadFile("./config.json")
	if fileErr != nil {
		fmt.Printf("File error: %v\n", fileErr)
		os.Exit(1)
	}

	var jsonType Config
	jsonErr := json.Unmarshal(file, &jsonType)
	if jsonErr != nil {
		fmt.Printf("Json parsing error: %v\n", jsonErr)
		os.Exit(1)
	}

	return jsonType
}
