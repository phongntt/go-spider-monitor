package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const CONF_FILE_PREFIX = "conf_"

func LoadAllConfig(confDirPath string) (map[string]ConfigData, error) {
	files, err := ioutil.ReadDir(confDirPath)
	if err != nil {
		return nil, err
	}

	allConfig := make(map[string]ConfigData)

	for _, f := range files {
		if !f.IsDir() && strings.HasPrefix(f.Name(), CONF_FILE_PREFIX) { // This is a file
			filePath := confDirPath + "/" + f.Name()
			confName := strings.TrimSuffix(strings.TrimPrefix(f.Name(), CONF_FILE_PREFIX), filepath.Ext(f.Name()))
			confData, err := ReadFromFile(filePath)
			if err != nil {
				allConfig[confName] = createErrorConfig()
			} else {
				allConfig[confName] = confData
			}
		}
	}

	return allConfig, nil
}

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
