package app

// All Locales supported by the app
var Locales = []Locale{"en", "sw"}

// Locale from which the descriptions and placeholders are going to be checked from
var mainLocale = "en"

// will look for all files ending with *_en.arb and *_sw.arb in the folders
// existing in this folder
const FROM = "/Users/mac/Programs/Personal/l10n_flutter/l10n"

// will look for app_en.arb and app_sw.arb here
const TO = "/Users/mac/Programs/Personal/l10n_flutter/l10n"

// Will attempt to run 'flutter pub get' command once the translations are combined
const FlutterProjectDir = "/Users/mac/Programs/Personal/l10n_flutter"

// Where the generated localization files will be stored
const LocalizationsDir = "/Users/mac/Programs/Personal/l10n_flutter/lib/localizations"
