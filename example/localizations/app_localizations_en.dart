import 'app_localizations.dart';

/// The translations for en).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  Account get account => AccountEn();

  @override
  AccountSettings get accountSettings => AccountSettingsEn();

  @override
  Profile get profile => ProfileEn();

  @override
  String greeting(String name) => "Hello $name,";
}

class AccountEn extends Account {
  @override
  String get information => "Account Information";

  @override
  String get settings => "Account Settings";
}

class AccountSettingsEn extends AccountSettings {
  @override
  String required(String label) => "$label is required";

  @override
  String get password => "Password";
}

class ProfileEn extends Profile {
  @override
  String get email => "Email";
}
