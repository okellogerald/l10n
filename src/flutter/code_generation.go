package flutter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/okellogerald/l10n.git/src/app"
	"github.com/okellogerald/l10n.git/src/utils"
	list "github.com/okellogerald/l10n.git/src/utils/list_utils"
)

func GenerateLocalizationFiles(settings app.GlobalSettings, data app.MethodsData) error {
	err := os.Mkdir(settings.To, 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%v/app_localizations.dart", settings.To))
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	w.WriteString(ln(`
	// ignore_for_file: depend_on_referenced_packages

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
	`))

	pred := func(s string) string {
		return fmt.Sprintf("Locale('%v')", s)
	}
	result := list.Map[string, string](app.Locales, pred)
	localesInstances := strings.Join(result, ",")
	w.WriteString(ln(fmt.Sprintf("static const List<Locale> supportedLocales = <Locale>[%v];", localesInstances)))
	w.WriteString(ln(""))

	writeInterfaceGroups(w, data.MethodGroups)
	writeInterfaceMethods(w, data.Methods)

	w.WriteString(ln("}"))

	w.WriteString(`
	class _AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {
		const _AppLocalizationsDelegate();
	  
		@override
		Future<AppLocalizations> load(Locale locale) {
		  return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
		}
			  
		@override
		bool shouldReload(_AppLocalizationsDelegate old) => false;
	`)

	locales := list.Map[string, string](app.Locales, func(s string) string {
		return fmt.Sprintf("'%v'", s)
	})

	w.WriteString(fmt.Sprintf(`
	@override
	bool isSupported(Locale locale) => <String>[%v].contains(locale.languageCode);
	`,
		strings.Join(locales, ","),
	),
	)

	w.WriteString(ln(`
	AppLocalizations lookupAppLocalizations(Locale locale) {
		// Lookup logic when only language code is specified.
		switch (locale.languageCode) {
	`))

	for i := 0; i < len(app.Locales); i++ {
		l1 := app.Locales[i]
		l2 := utils.CapitalizeFirstLetter(l1)
		s := fmt.Sprintf("case '%v': return AppLocalizations%v();", l1, l2)
		w.WriteString(ln(s))
	}

	w.WriteString(`
	}
	
	throw FlutterError(
		'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
		'an issue with the localizations generation tool. Please file an issue '
		'on GitHub with a reproducible sample app and the gen-l10n configuration '
		'that was used.'
	  );
	}
	}
	`)

	for i := 0; i < len(data.MethodGroups); i++ {
		group := data.MethodGroups[i]
		s := fmt.Sprintf("abstract class %v {", group.Identifier)
		w.WriteString(ln(s))
		w.WriteString(ln(""))

		writeInterfaceMethods(w, group.Methods)

		w.WriteString(ln("}"))
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = w.Flush()
	if err != nil {
		return err
	}

	err = generateSpecificLocaleFiles(settings, data)
	if err != nil {
		return err
	}

	log.Println("File writing completed.")

	return nil
}

func generateSpecificLocaleFiles(settings app.GlobalSettings, data app.MethodsData) error {
	for i := 0; i < len(app.Locales); i++ {
		locale := app.Locales[i]
		dir := settings.To
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

		writeMethodGroups(w, data.MethodGroups, locale)
		writeMethods(w, data.Methods, i)

		w.WriteString(ln("}"))
		w.WriteString(ln(""))

		for j := 0; j < len(data.MethodGroups); j++ {
			group := data.MethodGroups[j]
			s := fmt.Sprintf("class %[1]s%v extends %[1]s {", group.Identifier, capitalizedLocale)
			w.WriteString(ln(s))
			w.WriteString(ln(""))

			writeMethods(w, group.Methods, i)

			w.WriteString(ln("}"))
		}

		err = w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func writeInterfaceGroups(w *bufio.Writer, groups []app.MethodGroup) {
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		s := fmt.Sprintf("%v get %v;", group.Identifier, group.Name)
		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}
}

func writeInterfaceMethods(w *bufio.Writer, methods []app.Method) {
	for i := 0; i < len(methods); i++ {
		method := methods[i]
		hasPlaceHolders := len(method.PlaceHolders) > 0

		var s string

		if hasPlaceHolders {
			joiner := func(p app.PlaceHolder) string {
				return p.Type + " " + p.Name
			}
			result := list.Map[string, app.PlaceHolder](method.PlaceHolders, joiner)
			params := strings.Join(result, ", ")
			name := method.Name
			s = fmt.Sprintf("String %v(%v);", name, params)
		} else {
			s = fmt.Sprintf("String get %v;", method.Name)
		}

		var desc string
		if len(method.Description) == 0 {
			desc = fmt.Sprintf("No description provided for @%v", method.Name)
		} else {
			desc = method.Description
		}

		mainLocaleTranslation := method.Translations[0]
		mainLocaleTranslation = strings.ReplaceAll(mainLocaleTranslation, "\n", "")
		w.WriteString(lnt(fmt.Sprintf("/// %v", desc)))
		w.WriteString(lnt("///"))
		w.WriteString(lnt(fmt.Sprintf("/// In %v it is translated to:", app.MainLocale)))
		w.WriteString(lnt(fmt.Sprintf("/// **%v**", mainLocaleTranslation)))
		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}
}

func writeMethods(w *bufio.Writer, methods []app.Method, i int) {
	for k := 0; k < len(methods); k++ {
		method := methods[k]

		if len(method.Translations) < 2 {
			continue
		}

		w.WriteString(lnt("@override"))

		var s string
		tr := method.Translations[i]

		containsNewLines := strings.Contains(tr, "\n")
		hasPlaceHolders := len(method.PlaceHolders) > 0

		if hasPlaceHolders {
			joiner := func(p app.PlaceHolder) string {
				return p.Type + " " + p.Name
			}
			result := list.Map[string, app.PlaceHolder](method.PlaceHolders, joiner)
			params := strings.Join(result, ", ")
			value := replaceBrackets(tr)
			name := method.Name

			if containsNewLines {
				s = fmt.Sprintf("String %v(%v) => '''%v''';", name, params, value)
			} else {
				s = fmt.Sprintf("String %v(%v) => \"%v\";", name, params, value)
			}
		} else {
			if containsNewLines {
				s = fmt.Sprintf("String get %v => '''%v''';", method.Name, tr)
			} else {
				s = fmt.Sprintf("String get %v => \"%v\";", method.Name, tr)
			}
		}

		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}
}

func writeMethodGroups(w *bufio.Writer, groups []app.MethodGroup, locale string) {
	capitalizedLocale := utils.CapitalizeFirstLetter(locale)

	for i := 0; i < len(groups); i++ {
		group := groups[i]
		s := fmt.Sprintf("%[1]s get %[2]s => %[1]s%[3]s();", group.Identifier, group.Name, capitalizedLocale)
		w.WriteString(lnt("@override"))
		w.WriteString(lnt(s))
		w.WriteString(ln(""))
	}
}

func ln(s string) string {
	return fmt.Sprintf("%v\n", s)
}

// New line with tabs
func lnt(s string) string {
	return fmt.Sprintf("		%v\n", s)
}

func replaceBrackets(input string) string {
	re := regexp.MustCompile(`{([^}]+)}`)
	output := re.ReplaceAllString(input, `${$1}`)
	return output
}
