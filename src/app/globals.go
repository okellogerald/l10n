package app

var locales = []string{"en", "sw"}

// will look for all files ending with *_en.arb and *_sw.arb in the folders
// existing in this folder
const from = "/Users/mac/Programs/Frontend/lipa_link/l10n"

// will look for app_en.arb and app_sw.arb here
const to = "/Users/mac/Programs/Frontend/lipa_link/l10n"

// Will attempt to run 'flutter pub get' command once the translations are combined
const FlutterProjectDir = "/Users/mac/Programs/Frontend/lipa_link/"