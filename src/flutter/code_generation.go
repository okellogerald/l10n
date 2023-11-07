package flutter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/utils"
)

func GenerateLocalizationFiles(data app.MethodsData) error {
	err := os.Mkdir(app.LocalizationsDir, 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%v/app_localizations.dart", app.LocalizationsDir))
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered w
	w := bufio.NewWriter(file)

	// Write lines to the file

	// -> write starter logic
	w.WriteString(ln(`
	import 'dart:async';

	import 'package:flutter/foundation.dart';
	import 'package:flutter/widgets.dart';
	import 'package:flutter_localizations/flutter_localizations.dart';
	import 'package:intl/intl.dart' as intl;

	import 'app_localizations_en.dart';
	import 'app_localizations_sw.dart';

	abstract class AppLocalizations {
		AppLocalizations(String locale) : localeName = intl.Intl.canonicalizedLocale(locale.toString());
		
		final String localeName;

		static AppLocalizations? of(BuildContext context) {
	  		return Localizations.of<AppLocalizations>(context, AppLocalizations);
		}
  
		static const LocalizationsDelegate<AppLocalizations> delegate = _AppLocalizationsDelegate();

		static const List<LocalizationsDelegate<dynamic>> localizationsDelegates = <LocalizationsDelegate<dynamic>>[
			delegate,
			GlobalMaterialLocalizations.delegate,
			GlobalCupertinoLocalizations.delegate,
			GlobalWidgetsLocalizations.delegate,
		];

		static const List<Locale> supportedLocales = <Locale>[
		  Locale('en'),
		  Locale('sw')
		];
	`))

	for i := 0; i < len(data.MethodGroups); i++ {
		group := data.MethodGroups[i]
		s := fmt.Sprintf("%v get %v;", group.Identifier, group.Name)
		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}

	for i := 0; i < len(data.Methods); i++ {
		method := data.Methods[i]
		s := fmt.Sprintf("String get %v;", method.Name)
		var desc string
		if len(method.Description) == 0 {
			desc = fmt.Sprintf("/// No description provided for @%v", method.Name)
		} else {
			desc = method.Description
		}
		w.WriteString(lnt(desc))
		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}

	w.WriteString(ln("}"))

	w.WriteString(`
	class _AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {
		const _AppLocalizationsDelegate();
	  
		@override
		Future<AppLocalizations> load(Locale locale) {
		  return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
		}
	  
		@override
		bool isSupported(Locale locale) => <String>['en', 'sw'].contains(locale.languageCode);
	  
		@override
		bool shouldReload(_AppLocalizationsDelegate old) => false;
	  }
	  
	  AppLocalizations lookupAppLocalizations(Locale locale) {
	  
	  
		// Lookup logic when only language code is specified.
		switch (locale.languageCode) {
		  case 'en': return AppLocalizationsEn();
		  case 'sw': return AppLocalizationsSw();
		}
	  
		throw FlutterError(
		  'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
		  'an issue with the localizations generation tool. Please file an issue '
		  'on GitHub with a reproducible sample app and the gen-l10n configuration '
		  'that was used.'
		);
	  }
	  
	`)

	for i := 0; i < len(data.MethodGroups); i++ {
		group := data.MethodGroups[i]
		s := fmt.Sprintf("abstract class %v {", group.Identifier)
		w.WriteString(ln(s))
		w.WriteString(ln(""))

		for i := 0; i < len(group.Methods); i++ {
			method := data.Methods[i]
			s := fmt.Sprintf("String get %v;", method.Name)
			var desc string
			if len(method.Description) == 0 {
				desc = fmt.Sprintf("/// No description provided for @%v", method.Name)
			} else {
				desc = method.Description
			}
			w.WriteString(lnt(desc))
			w.WriteString(lnt(s))
			w.WriteString(ln(""))
		}

		w.WriteString(ln("}"))
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = w.Flush()
	if err != nil {
		return err
	}

	err = generateSpecificLocaleFiles(data)
	if err != nil {
		return err
	}

	log.Println("File writing completed.")

	return nil
}

func generateSpecificLocaleFiles(data app.MethodsData) error {
	for i := 0; i < len(app.Locales); i++ {
		locale := app.Locales[i]
		dir := app.LocalizationsDir
		file, err := os.Create(fmt.Sprintf("%v/app_localizations_%v.dart", dir, locale))
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a buffered w
		w := bufio.NewWriter(file)

		capitalizedLocale := utils.CapitalizeFirstLetter(locale)

		w.WriteString(ln("import 'app_localizations.dart';"))

		w.WriteString(ln(fmt.Sprintf(`
		/// The translations for %v).
		class AppLocalizations%v extends AppLocalizations {
  			AppLocalizations%v([String locale = '%v']) : super(locale);
		`, locale, capitalizedLocale, capitalizedLocale, locale)))

		for i := 0; i < len(data.MethodGroups); i++ {
			group := data.MethodGroups[i]
			s := fmt.Sprintf("%[1]s get %[2]s => %[1]s%[3]s();", group.Identifier, group.Name, capitalizedLocale)
			w.WriteString(lnt("@override"))
			w.WriteString(lnt(s))
			w.WriteString(ln(""))
		}

		for j := 0; j < len(data.Methods); j++ {
			method := data.Methods[j]
			w.WriteString(lnt("@override"))
			tr := method.Translations[i]

			containsNewLines := strings.Contains(tr, "\n")
			var s string
			if containsNewLines {
				s = fmt.Sprintf("String get %v => '''%v''';", method.Name, tr)
			} else {
				s = fmt.Sprintf("String get %v => \"%v\";", method.Name, tr)
			}
			w.WriteString(lnt(s))
			w.WriteString(ln(""))
		}

		w.WriteString(ln("}"))
		w.WriteString(ln(""))

		for j := 0; j < len(data.MethodGroups); j++ {
			group := data.MethodGroups[j]
			s := fmt.Sprintf("class %[1]s%v extends %[1]s {", group.Identifier, capitalizedLocale)
			w.WriteString(ln(s))
			w.WriteString(ln(""))

			for k := 0; k < len(group.Methods); k++ {
				method := data.Methods[k]
				w.WriteString(lnt("@override"))

				var s string
				tr := method.Translations[i]

				containsNewLines := strings.Contains(tr, "\n")

				if containsNewLines {
					s = fmt.Sprintf("String get %v => '''%v''';", method.Name, tr)
				} else {
					s = fmt.Sprintf("String get %v => \"%v\";", method.Name, tr)
				}
				w.WriteString(lnt(s))
				w.WriteString(ln(""))
			}

			w.WriteString(ln("}"))
		}

		err = w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func ln(s string) string {
	return fmt.Sprintf("%v\n", s)
}

// New line with tabs
func lnt(s string) string {
	return fmt.Sprintf("		%v\n", s)
}
