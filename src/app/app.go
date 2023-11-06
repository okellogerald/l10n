package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
)

func JoinTranslations() (*MethodsData, error) {
	/* err := os.RemoveAll(LocalizationsDir)
	if err != nil {
		return nil, err
	}
	*/
	folders, err := getAllFoldersFrom(FROM)
	if err != nil {
		return nil, err
	}

	var methods []Method
	var methodGroups []MethodGroup

	for i := 0; i < len(folders); i++ {
		hasAllLocales := checkIfFolderHasAllSpecifiedLocales(folders[i])
		if !hasAllLocales {
			notIncludedAllLocalesError := errors.New("Please make sure you have included all locales translations in all folders")
			if err != nil {
				return nil, notIncludedAllLocalesError
			}
		}

		files, err := getAllFilesFrom(folders[i])
		if err != nil {
			return nil, err
		}

		predicate := func(s string) bool {
			paths := strings.Split(s, "/")
			fileName := paths[len(paths)-1]
			return fileName == fmt.Sprintf("%v.json", mainLocale)
		}
		index := list.IndexWhere[string](files, predicate)
		if index == -1 {
			return nil, errors.New("could not find the main lcoale file")
		}

		mainFile := files[index]

		mainContent, err := DecodeJSONFile(mainFile)
		if err != nil {
			return nil, err
		}
		isGroup := IsMethodsGroup(mainContent)
		if isGroup {
			id, ok := mainContent["_id"].(string)
			if !ok {
				return nil, errors.ErrUnsupported
			}
			methodGroup, err := GetMethodsGroupFrom(id, files...)
			if err != nil {
				return nil, err
			}

			methodGroups = append(methodGroups, *methodGroup)
		} else {
			method, err := GetMethodsFrom(files...)
			if err != nil {
				return nil, err
			}

			methods = append(methods, method...)
		}
	}

	data := MethodsData{
		Methods:      methods,
		MethodGroups: methodGroups,
	}
	return &data, nil
}

func getMappedTranslationsFrom(b []byte) (Content, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(b), &data); err != nil {
		return nil, nil
	}

	return data, nil
}
