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

	name, ok := data["_name"].(string)
	if !ok {
		return ok
	}

	if len(identifier) == 0 || len(name) == 0 {
		return false
	}

	return true
}

func GetMethodsGroupFrom(identifier string, name string, files ...string) (*MethodGroup, error) {
	methods, err := GetMethodsFrom(files...)
	if err != nil {
		return nil, err
	}
	group := MethodGroup{
		Identifier: identifier,
		Name:       name,
		Methods:    methods,
	}
	return &group, nil
}

func GetMethodsFrom(files ...string) ([]Method, error) {
	var data = make(map[string]interface{})

	for i := 0; i < len(files); i++ {
		content, err := DecodeJSONFile(files[i])
		if err != nil {
			return nil, err
		}

		locale, err := getLocaleFromFile(files[i])
		if err != nil {
			return nil, err
		}

		data[locale] = content
	}

	var methods []Method
	mainLocaleContent, ok := utils.ConvertTo[map[string]interface{}](data[MainLocale], data)
	if !ok {
		return nil, errors.New("Main locale content - something is wrong")
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
			translation, ok := utils.ConvertTo[map[string]interface{}](data[Locales[i]], data)
			if !ok {
				return nil, errors.New("Locale translation - Something went wrong")
			}
			if translation[name] == nil {
				continue
			}
			value, ok := translation[name].(string)

			if !ok {
				return nil, errors.New("Value - Something went wrong")
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
