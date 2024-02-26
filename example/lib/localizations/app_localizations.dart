
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
	
static const List<Locale> supportedLocales = <Locale>[Locale('en'),Locale('sw')];

		ProfilePage get profilePage;

		ProfileEditPage get profileEditPage;

		/// A greeting
		///
		/// In en it is translated to:
		/// **Hello {name},**
		String greeting(String name);

}

	class _AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {
		const _AppLocalizationsDelegate();
	  
		@override
		Future<AppLocalizations> load(Locale locale) {
		  return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
		}
			  
		@override
		bool shouldReload(_AppLocalizationsDelegate old) => false;
	
	@override
	bool isSupported(Locale locale) => <String>['en','sw'].contains(locale.languageCode);
	
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
	}
	abstract class ProfilePage {

		/// No description provided for @hello
		///
		/// In en it is translated to:
		/// **Hello**
		String get hello;

}
abstract class ProfileEditPage {

		/// No description provided for @edit
		///
		/// In en it is translated to:
		/// **Edit Profile**
		String get edit;

}
