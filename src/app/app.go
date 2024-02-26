package app

import (
	"errors"
	"fmt"
	"strings"

	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
)

func JoinTranslations(settings GlobalSettings) (*MethodsData, error) {
	folders, err := getAllFoldersFrom(settings.From)
	fmt.Println(folders)
	if err != nil {
		return nil, err
	}

	var methods []Method
	var methodGroups []MethodGroup

	for i := 0; i < len(folders); i++ {
		hasAllLocales := checkIfFolderHasAllSpecifiedLocales(folders[i])
		if !hasAllLocales {
			notIncludedAllLocalesError := errors.New("please make sure you have included all locales translations in all folders")
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
			return fileName == fmt.Sprintf("%v.json", MainLocale)
		}
		index := list.IndexWhere[string](files, predicate)
		if index == -1 {
			return nil, errors.New("could not find the main lcoale file")
		}

		mainFile := files[index]

		mainContent, mainContentList, err := DecodeJSONFile(mainFile)
		if err != nil {
			return nil, err
		}
		if mainContent != nil {
			isGroup := IsMethodsGroup(mainContent)
			if isGroup {
				id := mainContent["_id"].(string)
				name := mainContent["_name"].(string)
				methodGroup, err := GetMethodsGroupFrom(id, name, files...)
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

		if mainContentList != nil {
			groups, err := GetMethodsGroupsFrom(files...)
			if err != nil {
				return nil, err
			}

			methodGroups = append(methodGroups, groups...)
		}
	}

	data := MethodsData{
		Methods:      methods,
		MethodGroups: methodGroups,
	}
	return &data, nil
}
