package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFromFile(filename string) (ConfigData, error) {
	var config ConfigData

	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return config, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, err
	}

	err = config.DecryptAll()
	if err != nil {
		var emptyConfig ConfigData
		return emptyConfig, err
	}

	return config, nil
}
