package jsonmanager

import (
	"encoding/json"
	"os"
)

type JsonManager struct {
	filePath string
}

func New(path string) JsonManager {
	return JsonManager{filePath: path}
}

func (f JsonManager) ReadLines() ([]byte, error) {
	return os.ReadFile(f.filePath)
}

func (f JsonManager) WriteResult(data any) error {
	jsonString, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, jsonString, 0644)
}
