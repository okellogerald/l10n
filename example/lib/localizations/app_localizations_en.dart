import 'app_localizations.dart';

/// The translations for en).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  Profile get profile => ProfileEn();

  @override
  String greeting(String name) => "Hello $name,";
}

class ProfileEn extends Profile {
  @override
  String get email => "Email";
}
