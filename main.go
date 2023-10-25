package main

import (
	app "github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/flutter"
)

func main() {
	app.JoinTranslations()
	flutter.PubGet()
}
