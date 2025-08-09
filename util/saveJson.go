package util

import (
	"encoding/json"
	"log"
	"os"
)

func SaveJson(savingJson map[string]interface{}, path string) {
	data, err := json.Marshal(savingJson)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}
}
