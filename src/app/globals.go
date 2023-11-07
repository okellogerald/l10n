package app

// All Locales supported by the app
var Locales = []Locale{"en", "sw"}

// Locale from which the descriptions and placeholders are going to be checked from
var MainLocale = "en"

// will look for all files ending with *_en.json and *_sw.json in the folders
// existing in this folder
const FROM = "/Users/mac/Programs/Personal/l10n/example/l10n"

// Will attempt to run 'flutter pub get' command once the translations are combined
const FlutterProjectDir = "/Users/mac/Programs/Personal/l10n/example/"

// Where the generated localization files will be stored
const LocalizationsDir = "/Users/mac/Programs/Personal/l10n/example/lib/localizations"
