import 'app_localizations.dart';

/// The translations for sw).
class AppLocalizationsSw extends AppLocalizations {
  AppLocalizationsSw([String locale = 'sw']) : super(locale);

  @override
  Account get account => AccountSw();

  @override
  AccountSettings get accountSettings => AccountSettingsSw();

  @override
  Profile get profile => ProfileSw();

  @override
  String greeting(String name) => "Habari $name";
}

class AccountSw extends Account {
  @override
  String get information => "Taarifa za Akaunti";

  @override
  String get settings => "Mipangilio ya Akaunti";
}

class AccountSettingsSw extends AccountSettings {
  @override
  String required(String label) => "$label is required";

  @override
  String get password => "Nenosiri";
}

class ProfileSw extends Profile {
  @override
  String get email => "Barua Pepe";
}
