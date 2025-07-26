package init

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
