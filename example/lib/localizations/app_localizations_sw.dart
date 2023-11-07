import 'app_localizations.dart';

/// The translations for sw).
class AppLocalizationsSw extends AppLocalizations {
  AppLocalizationsSw([String locale = 'sw']) : super(locale);

  @override
  Profile get profile => ProfileSw();

  @override
  String greeting(String name) => "Habari $name";
}

class ProfileSw extends Profile {
  @override
  String get email => "Barua Pepe";
}
