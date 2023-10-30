package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	maputils "github.com/okellogerald/l10n.git/src/utils/map_utils"
)

func JoinTranslations() error {
	for i := 0; i < len(locales); i++ {
		path := fmt.Sprintf("%v/app_%v.arb", from, locales[i])
		println(path)
		err := deleteFileContents(path)
		if err != nil {
			return err
		}
	}

	folders, err := getAllFoldersFrom(from)
	if err != nil {
		return err
	}

	// root folder
	hasAllLocales := checkIfRootFolderHasAllSpecifiedLocales(from)
	if !hasAllLocales {
		notIncludedAllLocalesError := errors.New("Please make sure you have included all locales translations in root folders")
		if err != nil {
			return notIncludedAllLocalesError
		}
	}

	for i := 0; i < len(folders); i++ {
		hasAllLocales := checkIfFolderHasAllSpecifiedLocales(folders[i])
		if !hasAllLocales {
			notIncludedAllLocalesError := errors.New("Please make sure you have included all locales translations in all folders")
			if err != nil {
				return notIncludedAllLocalesError
			}
		}

		files, err := getAllFilesFrom(folders[i])
		if err != nil {
			return err
		}

		for i := 0; i < len(files); i++ {
			content, err := readARBFile(files[i])
			if err != nil {
				return err
			}

			locale, err := getLocaleFromFile(files[i])
			if err != nil {
				return err
			}

			rootContentPath := fmt.Sprintf("%v/app_%v.arb", from, locale)
			rootContent, err := readARBFile(rootContentPath)
			if err != nil {
				return err
			}

			mergedData, err := mergeContents(rootContent, content)
			if err != nil {
				return err
			}
			
			to := fmt.Sprintf("%v/app_%v.arb", from, locale)
			if err := writeARB(mergedData, to); err != nil {
				return err
			}
		}
	}

	return nil
}

func mergeContents(b1, b2 []byte) (ARBData, error) {
	data1, err := getMappedTranslationsFrom(b1)
	data2, err := getMappedTranslationsFrom(b2)
	if err != nil {
		return nil, err
	}

	mergedData := maputils.Combine[string, interface{}](data1, data2)
	return mergedData, nil
}

func getMappedTranslationsFrom(b []byte) (ARBData, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(b), &data); err != nil {
		return nil, nil
	}

	return data, nil
}

func deleteFileContents(filePath string) error {
	// Open the file in write mode
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Truncate the file's contents
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	file.WriteString(`
	{
		
	}
	`)

	fmt.Println("File contents deleted.")
	return nil
}
