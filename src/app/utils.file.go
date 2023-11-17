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

func DecodeJSONFile(filePath string) (Content, ContentList, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err == nil {
		return m, nil, err
	}

	var l []map[string]interface{}
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, nil, err
	}

	return nil, l, nil
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
		return "", errors.New("please check the filenames")
	}

	index := list.IndexOf[string](locale, Locales)
	if index == -1 {
		return "", errors.New("file not found")
	}

	return Locales[index], nil
}
