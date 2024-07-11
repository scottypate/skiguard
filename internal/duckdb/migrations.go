package duckdb

import (
	"fmt"
	"log/slog"
	"os"
)

func RunMigrations(dir string) {
	files, err := os.ReadDir(dir)

	if err != nil {
		slog.Error(fmt.Sprintf("Error reading directory: %s", err.Error()))
	}

	for _, file := range files {
		fileByte, err := os.ReadFile(dir + "/" + file.Name())
		if err != nil {
			slog.Error(fmt.Sprintf("Error reading file: %s", err.Error()))
		}
		Execute(string(fileByte))
	}
}
