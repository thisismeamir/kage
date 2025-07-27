package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveForm(path string, form interface{}) error {
	data, err := json.MarshalIndent(form, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
