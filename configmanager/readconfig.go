package configmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var configFile = fmt.Sprintf("%v/src/github.com/deeper-x/gocamole/config/global.json", os.Getenv("GOPATH"))

// ConfigFormat configuration file skeleton
type ConfigFormat struct {
	Name string `json:"name"`
}

var cf ConfigFormat

// ConfigFileExists check is path exists
func ConfigFileExists() bool {
	_, err := os.Stat(configFile)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// OpenConfigFile os open wrapper
func OpenConfigFile() *os.File {
	jsonfile, err := os.Open(configFile)

	if err != nil {
		log.Fatal(err)
	}

	return jsonfile
}

// LoadConfigFile load configuration file
func LoadConfigFile() []byte {
	jsonfile := OpenConfigFile()

	byteValue, err := ioutil.ReadAll(jsonfile)

	defer jsonfile.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(byteValue, &cf); err != nil {
		log.Fatal("unmarshal error")
	}

	return byteValue
}

// ReadConfig given a key, returns its value in configuration
func ReadConfig() ConfigFormat {
	fileContent := LoadConfigFile()

	json.Unmarshal(fileContent, &cf)

	return cf
}
