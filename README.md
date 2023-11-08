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
1. Create translation folders. They should all be in the same folder.
   Example:
        profile_page
            en.json
            ru.json
2. Go into `src/app/globals.go` and change the constants there to match your projects details
3. Then run the program with the command `go run .`. It should generate all the localization files you need, just as `flutter_localizations` does it and more.

The `example/` folder may be helpful.