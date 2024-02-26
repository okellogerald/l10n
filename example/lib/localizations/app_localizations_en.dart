import 'app_localizations.dart';

		/// The translations for en).
		class AppLocalizationsEn extends AppLocalizations {
  			AppLocalizationsEn([String locale = 'en']) : super(locale);
		
		@override
		ProfilePage get profilePage => ProfilePageEn();

		@override
		ProfileEditPage get profileEditPage => ProfileEditPageEn();

		@override
		String greeting(String name) => "Hello $name,";

}

class ProfilePageEn extends ProfilePage {

		@override
		String get hello => "Hello";

}
class ProfileEditPageEn extends ProfileEditPage {

		@override
		String get edit => "Edit Profile";

}
