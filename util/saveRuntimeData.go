package util

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	language_handlers "github.com/thisismeamir/kage/internal/engine/execution-system/language-handlers"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
)

// SaveRuntimeDataInCSV appends ExecutionResult data to a CSV file.
func SaveRuntimeDataInCSV(result language_handlers.ExecutionResult, executionIdentifier string, conf config.Config) {
	savingPath := conf.BasePath + "/data/runtime/runtime-statistics.csv"
	os.MkdirAll(filepath.Dir(savingPath), os.ModePerm)

	_, err := os.Stat(savingPath)
	fileExists := err == nil

	file, err := os.OpenFile(savingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExists {
		writer.Write([]string{
			"ExecutionID", "Error", "ExitCode", "Duration", "OutputJsonPath", "Stdout", "Stderr",
		})
	}

	writer.Write([]string{
		executionIdentifier,
		result.Error,
		strconv.Itoa(result.ExitCode),
		result.Duration.String(),
		result.OutputJsonPath,
		result.Stdout,
		result.Stderr,
	})
}
