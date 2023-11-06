package main

import (
	"fmt"

	"github.com/okellogerald/l10n.git/src/app"
)

func main() {
	data, err := app.JoinTranslations()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
