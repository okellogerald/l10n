package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
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

			var parsedData1 map[string]interface{}
			err = json.Unmarshal([]byte(rootContent), &parsedData1)
			if err != nil {
				return err
			}

			var parsedData2 map[string]interface{}
			err = json.Unmarshal([]byte(content), &parsedData2)
			if err != nil {
				return err
			}

			list1 := convertMapToARBData(parsedData1)
			list2 := convertMapToARBData(parsedData2)
			mergedData := list.Combine(list1, list2)

			to := fmt.Sprintf("%v/app_%v.arb", from, locale)
			err = writeARB(mergedData, to)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
