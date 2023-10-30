package main

import (
	"fmt"

	"github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/flutter"
)

func main() {
	path := fmt.Sprintf("%v/app_%v.arb", app.FROM, "en")
	data, err := app.ReadARBFile(path)
	if err != nil {
		panic(err)
	}

	list := flutter.GetMethodsFrom(data)

	for i := 0; i < len(list); i++ {
		method := list[i]
		if method.Name != "greeting" {
			continue
		}
		e := fmt.Sprintf(`
		Method:
			name: %v
			desc: %v
			placeholders: %v
		`,
			method.Name,
			method.Description,
			method.PlaceHolders,
		)
		println(e)
	}
	flutter.GenerateAppLocalizationFile()
	/* err := app.JoinTranslations()
	if err != nil {
		panic(err)
	}

	flutter.PubGet() */
}
