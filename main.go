package main

import (
	"os"

	"github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/flutter"
)

func main() {
	data, err := app.JoinTranslations()
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll(app.LocalizationsDir)
	if err != nil {
		panic(err)
	}

	err = flutter.GenerateLocalizationFiles(*data)
	if err != nil {
		panic(err)
	}

	err = flutter.PubGet()
	if err != nil {
		panic(err)
	}

	err = flutter.Format()
	if err != nil {
		panic(err)
	}
}
