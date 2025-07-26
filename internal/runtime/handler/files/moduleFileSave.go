package files

import (
	"encoding/json"
	module "github.com/thisismeamir/kage/pkg/module"
	"os"
)

func ModuleFileSave(path string, atomModel module.ModuleModel) error {
	file, err := os.Create(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, create it
			file, err = os.Create(path)
			if err != nil {
				panic("failed to create file: " + err.Error())
			}
		} else {
			panic("failed to open file: " + err.Error())
		}
	}
	defer file.Close()

	data, err := json.MarshalIndent(atomModel, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
