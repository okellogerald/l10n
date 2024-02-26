import 'app_localizations.dart';

		/// The translations for sw).
		class AppLocalizationsSw extends AppLocalizations {
  			AppLocalizationsSw([String locale = 'sw']) : super(locale);
		
		@override
		ProfilePage get profilePage => ProfilePageSw();

		@override
		ProfileEditPage get profileEditPage => ProfileEditPageSw();

		@override
		String greeting(String name) => "Habari $name";

}

class ProfilePageSw extends ProfilePage {

		@override
		String get hello => "Habari";

}
class ProfileEditPageSw extends ProfileEditPage {

		@override
		String get edit => "Badili Wasifu";

}
