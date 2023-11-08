package app

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
)

func getAllFoldersFrom(path string) ([]string, error) {
	var folders []string

	// Read the contents of the folder
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Iterate over the entries and identify folders
	for _, entry := range entries {
		if entry.IsDir() {
			folderName := entry.Name()
			folderPath := filepath.Join(path, folderName)
			folders = append(folders, folderPath)
		}
	}

	return folders, nil
}

func getAllFilesFrom(folderPath string) ([]string, error) {
	var files []string

	// Read the contents of the folder
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	// Iterate over the entries and identify folders
	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			path := filepath.Join(folderPath, name)
			files = append(files, path)
		}
	}

	return files, nil
}

func checkIfFolderHasAllSpecifiedLocales(folderPath string) bool {
	files, err := getAllFilesFrom(folderPath)
	if err != nil {
		return false
	}

	if len(files) != len(Locales) {
		return false
	}

	for i := 0; i < len(Locales); i++ {
		fileName := path.Base(files[i])
		fileName, found := strings.CutSuffix(fileName, ".json")
		contanined := list.CheckFor[string](fileName, Locales)
		if !contanined || !found {
			return false
		}
	}

	return true
}

func checkIfRootFolderHasAllSpecifiedLocales(folderPath string) bool {
	files, err := getAllFilesFrom(folderPath)
	if err != nil {
		return false
	}

	if len(files) != len(Locales) {
		return false
	}

	for i := 0; i < len(Locales); i++ {
		fileName := path.Base(files[i])
		fileName, found := strings.CutSuffix(fileName, ".json")
		fileName, found = strings.CutPrefix(fileName, "app_")
		contanined := list.CheckFor[string](fileName, Locales)
		if !contanined || !found {
			return false
		}
	}

	return true
}

func DecodeJSONFile(filePath string) (Content, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func WriteJSONFile(filename string, data Content) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func getLocaleFromFile(filePath string) (string, error) {
	fileName := path.Base(filePath)
	locale, found := strings.CutSuffix(fileName, ".json")
	if !found {
		return "", errors.New("Please check the filenames")
	}

	index := list.IndexOf[string](locale, Locales)
	if index == -1 {
		return "", errors.New("File not found")
	}

	return Locales[index], nil
}

func writeARB(data Content, filePath string) error {
	// Convert the ARBData map to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to an ARB file
	os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
