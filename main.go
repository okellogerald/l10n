package main

import (
	"fmt"
	"os"

	"github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/flutter"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	settings := app.GlobalSettings {
        FlutterProjectDir:   dir,
		LocalizationsDir: fmt.Sprintf("%v/l10n", dir),
        LocalizationsOutDir:   fmt.Sprintf("%v/lib/localizations/", dir),
	}

	data, err := app.JoinTranslations(settings)
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll(settings.LocalizationsDir)
	if err != nil {
		panic(err)
	}

	err = flutter.GenerateLocalizationFiles(settings, *data)
	if err != nil {
		panic(err)
	}

	err = flutter.PubGet(settings)
	if err != nil {
		panic(err)
	}

	err = flutter.Format(settings)
	if err != nil {
		panic(err)
	}

	err = flutter.ApplyFixes(settings)
	if err != nil {
		panic(err)
	}
}
