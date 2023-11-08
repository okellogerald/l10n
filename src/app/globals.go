package app

// All Locales supported by the app
var Locales = []Locale{"en", "sw"}

// Locale from which the descriptions and placeholders are going to be checked from
const MainLocale = "en"

// will look for localization files in this folder
const FROM = "/Users/mac/Programs/Personal/l10n/example/l10n"

// Will attempt to run 'flutter pub get' command once the translations are combined
const FlutterProjectDir = "/Users/mac/Programs/Personal/l10n/example/"

// Where the generated localization files will be stored
const LocalizationsDir = "/Users/mac/Programs/Personal/l10n/example/lib/localizations"

