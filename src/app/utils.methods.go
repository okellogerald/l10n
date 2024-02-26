package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/okellogerald/l10n.git/src/utils"
	maputils "github.com/okellogerald/l10n.git/src/utils/map_utils"
)

func IsMethodsGroup(data map[string]interface{}) bool {
	identifier, ok := data["_id"].(string)
	if !ok {
		return ok
	}

	if len(identifier) == 0 {
		return false
	}

	return true
}

func GetMethodsGroupFrom(identifier string, files ...string) (*MethodGroup, error) {
	methods, err := GetMethodsFrom(files...)
	if err != nil {
		return nil, err
	}
	group := MethodGroup{
		Identifier: identifier,
		Methods:    methods,
	}
	return &group, nil
}

func GetMethodsFrom(files ...string) ([]Method, error) {
	var localesMap = make(map[string]interface{})

	for i := 0; i < len(files); i++ {
		content, contentList, err := DecodeJSONFile(files[i])
		if err != nil {
			return nil, err
		}

		locale, err := getLocaleFromFile(files[i])
		if err != nil {
			return nil, err
		}

		if content != nil {
			localesMap[locale] = content
		}

		if contentList != nil {
			return nil, errors.New("unexpected list here")
		}
	}

	var methods []Method
	mainLocaleContent, ok := utils.ConvertTo[map[string]interface{}](localesMap[MainLocale], localesMap)
	if !ok {
		return nil, errors.New("main locale content - something is wrong")
	}

	for k := range mainLocaleContent {
		name := k
		var placeHolders []PlaceHolder
		var desc string

		if name == "_id" || name == "_name" {
			continue
		}

		if ok := strings.HasPrefix(k, "@"); ok {
			continue
		}

		more := mainLocaleContent[fmt.Sprintf("@%v", name)]

		map1, ok := maputils.ConvertToContentMap(more)
		if ok && len(map1) != 0 {
			place_holders := map1["placeholders"]
			description := map1["description"]

			if map2, ok := place_holders.(map[string]interface{}); ok {
				for k, v := range map2 {
					if map3, ok := v.(map[string]interface{}); ok {
						_type := fmt.Sprintf("%v", map3["type"])
						placeHolder := PlaceHolder{
							Name: k,
							Type: _type,
						}
						placeHolders = append(placeHolders, placeHolder)
					} else {
						println("2n not of a good placeholders type")
					}
				}
			} else {
				println("not of a good placeholders type")
			}

			if v2, ok := description.(string); ok {
				desc = v2
			}
		}

		var translations []string
		for i := 0; i < len(Locales); i++ {
			translation, ok := utils.ConvertTo[map[string]interface{}](localesMap[Locales[i]], localesMap)
			if !ok {
				return nil, errors.New("Locale translation - Something went wrong")
			}
			if translation[name] == nil {
				continue
			}
			value, ok := translation[name].(string)
			if !ok {
				return nil, errors.New("value - Something went wrong")
			}

			translations = append(translations, value)
		}

		method := Method{
			Name:         name,
			Description:  desc,
			PlaceHolders: placeHolders,
			Translations: translations,
		}
		methods = append(methods, method)
	}

	return methods, nil
}

type partialGroup struct {
	id      string
	name    string
	methods []partialMethod
}

type partialMethod struct {
	name         string
	description  string
	placeHolders []PlaceHolder
}

// Responsible for transcoding files with a list of method-groups
func GetMethodsGroupsFrom(files ...string) ([]MethodGroup, error) {
	// detecting the main file
	mainFile := ""
	for i := 0; i < len(files); i++ {
		locale, err := getLocaleFromFile(files[i])
		if err != nil {
			return nil, err
		}

		if locale == MainLocale {
			mainFile = files[i]
			break
		}
	}

	// getting main-content from the main-locale file
	_, mainContentList, err := DecodeJSONFile(mainFile)
	if err != nil {
		return nil, err
	}

	// getting groups information
	var partialGroups = make([]partialGroup, 0)

	for i := 0; i < len(mainContentList); i++ {
		item := mainContentList[i]
		isGroup := IsMethodsGroup(item)
		if isGroup {
			id := item["_id"].(string)
			methods := make([]partialMethod, 0)

			for k := range item {
				if k == "_id" || k == "_name" {
					continue
				}
				if strings.HasPrefix(k, "@") {
					continue
				}

				methodName := k
				var methodDesc string
				var methodPlaceHolders []PlaceHolder

				more := item[fmt.Sprintf("@%v", methodName)]

				map1, ok := maputils.ConvertToContentMap(more)
				if ok && len(map1) != 0 {
					placeholders := map1["placeholders"]
					description := map1["description"]

					if map2, ok := placeholders.(map[string]interface{}); ok {
						for k, v := range map2 {
							if map3, ok := v.(map[string]interface{}); ok {
								_type := fmt.Sprintf("%v", map3["type"])
								p := PlaceHolder{Name: k, Type: _type}
								methodPlaceHolders = append(methodPlaceHolders, p)
							}
						}
					}
					if v2, ok := description.(string); ok {
						methodDesc = v2
					}
				}
				method := partialMethod{
					name:         methodName,
					description:  methodDesc,
					placeHolders: methodPlaceHolders,
				}
				methods = append(methods, method)
			}

			group := partialGroup{
				id:      id,
				methods: methods,
			}
			partialGroups = append(partialGroups, group)
		} else {
			return nil, errors.New("expecting only method groups")
		}
	}

	// converting list of translations to maps
	var localeTranslations = make(map[string]map[string]map[string]string)
	for i := 0; i < len(files); i++ {
		locale, err := getLocaleFromFile(files[i])
		if err != nil {
			return nil, err
		}

		var contentList ContentList

		if locale == MainLocale {
			contentList = mainContentList
		} else {
			_, list, err := DecodeJSONFile(files[i])
			if err != nil {
				return nil, err
			}
			contentList = list
		}

		var groupContentMap = make(map[string]map[string]string)
		for j := 0; j < len(contentList); j++ {
			partialGroup := partialGroups[j]
			content := contentList[j]

			newContent := make(map[string]string)
			for k, v := range content {
				if k == "_id" || k == "_name" {
					continue
				}
				if strings.HasPrefix(k, "@") {
					continue
				}

				if value, ok := v.(string); ok {
					newContent[k] = value
				}
			}
			groupContentMap[partialGroup.id] = newContent
		}

		localeTranslations[locale] = groupContentMap
	}

	// at last creating method groups
	var groups = make([]MethodGroup, 0)
	for i := 0; i < len(partialGroups); i++ {
		g := partialGroups[i]
		var methods = make([]Method, 0)
		for j := 0; j < len(g.methods); j++ {
			var translations = make([]string, 0)
			for k := 0; k < len(Locales); k++ {
				translation := localeTranslations[Locales[k]][g.id][g.methods[j].name]
				translations = append(translations, translation)
			}

			m := g.methods[j]

			method := Method{
				Name:         m.name,
				Description:  m.description,
				PlaceHolders: m.placeHolders,
				Translations: translations,
			}
			methods = append(methods, method)
		}

		group := MethodGroup{
			Identifier: g.id,
			Methods:    methods,
		}
		groups = append(groups, group)
	}

	return groups, nil
}
