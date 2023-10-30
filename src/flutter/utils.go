package flutter

import (
	"fmt"
	"strings"
)

type GotComments = map[string]interface{}

func GetMethodsFrom(data map[string]interface{}) []Method {
	var methods []Method

	for k := range data {
		var placeHolders []PlaceHolder
		var desc string

		if ok := strings.HasPrefix(k, "@"); ok {
			continue
		}

		more := data[fmt.Sprintf("@%v", k)]

		if map1, ok := more.(GotComments); ok {
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

		method := Method{
			Name:         k,
			Description:  desc,
			PlaceHolders: placeHolders,
		}
		methods = append(methods, method)
	}

	return methods
}
