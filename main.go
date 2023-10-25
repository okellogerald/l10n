package main

import (
	app "github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/flutter"
)

func main() {
	err := app.JoinTranslations()
	if err != nil {
		panic(err)
	}

	flutter.PubGet()
}
