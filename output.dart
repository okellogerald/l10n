
	abstract class AppLocalizations
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
	
}

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
	  
	