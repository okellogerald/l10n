package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
)

func main() {
	deleteSampleFiles()

	folders, err := getAllFoldersFrom(from)
	handleError(err)

	// root folder
	hasAllLocales := checkIfRootFolderHasAllSpecifiedLocales(from)
	if !hasAllLocales {
		notIncludedAllLocalesError := errors.New("Please make sure you have included all locales translations in root folders")
		handleError(notIncludedAllLocalesError)
	}

	for i := 0; i < len(folders); i++ {
		hasAllLocales := checkIfFolderHasAllSpecifiedLocales(folders[i])
		if !hasAllLocales {
			notIncludedAllLocalesError := errors.New("Please make sure you have included all locales translations in all folders")
			handleError(notIncludedAllLocalesError)
		}

		files, err := getAllFilesFrom(folders[i])
		handleError(err)

		for i := 0; i < len(files); i++ {
			content, err := readARBFile(files[i])
			handleError(err)

			locale, err := getLocaleFromFile(files[i])
			handleError(err)

			rootContentPath := fmt.Sprintf("%v/app_%v.arb", from, locale)
			rootContent, err := readARBFile(rootContentPath)
			handleError(err)

			var parsedData1 map[string]interface{}
			err = json.Unmarshal([]byte(rootContent), &parsedData1)
			if err != nil {
				fmt.Println("Error parsing ARB file 1:", err)
				return
			}

			var parsedData2 map[string]interface{}
			err = json.Unmarshal([]byte(content), &parsedData2)
			if err != nil {
				fmt.Println("Error parsing ARB file 2:", err)
				return
			}

			list1 := convertMapToARBData(parsedData1)
			list2 := convertMapToARBData(parsedData2)
			mergedData := list.Combine(list1, list2)

			to := fmt.Sprintf("%v/app_%v.arb", from, locale)
			err = writeARB(mergedData, to)
			handleError(err)
		}
	}
}

func deleteSampleFiles() {
	sampleFile1 := fmt.Sprintf("%v/sample_output_en.arb", from)
	sampleFile2 := fmt.Sprintf("%v/sample_output_sw.arb", from)

	os.Remove(sampleFile1)
	os.Remove(sampleFile2)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
