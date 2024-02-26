package app

// All Locales supported by the app
var Locales = []Locale{"en", "sw"}

// Locale from which the descriptions and placeholders are going to be checked from
const MainLocale = "en"

type GlobalSettings struct {
	// will look for localization files in this folder
	From string

	// Where the generated localization files will be stored
	To string

	FlutterProjectDir string
}
