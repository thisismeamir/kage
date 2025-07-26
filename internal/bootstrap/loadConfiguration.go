package bootstrap

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func LoadConfiguration(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty Config
			log.Println("file does not exist:", path)
			return Config{}
		}
		// If there is another error, log it and return an empty Config
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}
	}
	return cfg
}

var configPath = "./configs/default.conf.json"
var globalConfig Config

func SetConfigPath(path string) {
	configPath = path
}

func GetConfigPath() string {
	return configPath
}

func SetGlobalConfig(cfg Config) {
	globalConfig = cfg
}

func GetGlobalConfig() Config {
	return globalConfig
}

func SaveConfigFile() error {
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(globalConfig, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
