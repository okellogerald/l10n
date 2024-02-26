Making it possible to divide translations into subfolders instead of having one big translation file.

If [this issue](https://github.com/flutter/flutter/issues/107157) is resolved, then that would be a preferred solution.

## Requirements
1. Go should be installed in your system
2. The following dependencies should be added in pubspec.yaml
   -   intl:
   -   flutter_localizations:
            sdk: flutter
   Also make sure that the l10n.yaml file is deleted, because its presence will prompt generating the localization files with the flutter_localizations package.

## Getting Started
1. Run `go build -o {name_of_the_executable e.g locgen}`
2. Run `mv locgen /usr/local/bin`
3. Run `locgen` at the root of your flutter project

The `example/` folder may be helpful.