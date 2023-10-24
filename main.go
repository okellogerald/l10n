package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
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


			println("--------------------")
			println("before merge:")
			println(string(content))
			println(string(rootContent))

			var parsedData1 ARBData
			err = json.Unmarshal([]byte(content), &parsedData1)
			if err != nil {
				fmt.Println("Error parsing ARB file 1:", err)
				return
			}

			var parsedData2 ARBData
			err = json.Unmarshal([]byte(rootContent), &parsedData2)
			if err != nil {
				fmt.Println("Error parsing ARB file 2:", err)
				return
			}

			mergedData := mergeARBData(parsedData1, parsedData2)

			println("--------------------")
			println("after merge:")
			println(fmt.Sprintf("%v", mergedData))
			println("--------------------")

			to := fmt.Sprintf("%v/sample_output_%v.arb", from, locale)
			err = writeARB(mergedData, to)
			handleError(err)
		}
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
