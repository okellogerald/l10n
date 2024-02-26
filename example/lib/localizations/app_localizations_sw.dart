import 'app_localizations.dart';

/// The translations for sw).
class AppLocalizationsSw extends AppLocalizations {
  AppLocalizationsSw([String locale = 'sw']) : super(locale);

  @override
  RequestCard get requestCard => RequestCardSw();

  @override
  GotYourRequest get gotYourRequest => GotYourRequestSw();

  @override
  FundTopUp get fundTopUp => FundTopUpSw();

  @override
  FundTopUpByPhone get fundTopUpByPhone => FundTopUpByPhoneSw();

  @override
  FundTopUpByBank get fundTopUpByBank => FundTopUpByBankSw();

  @override
  Balance get balance => BalanceSw();

  @override
  QuickActions get quickActions => QuickActionsSw();

  @override
  HomeCards get homeCards => HomeCardsSw();

  @override
  Profile get profile => ProfileSw();

  @override
  ProfilePage get profilePage => ProfilePageSw();

  @override
  SplashPage get splashPage => SplashPageSw();

  @override
  Transactions get transactions => TransactionsSw();

  @override
  Transaction get transaction => TransactionSw();

  @override
  TransferType get transferType => TransferTypeSw();

  @override
  TransferToBankPage get transferToBankPage => TransferToBankPageSw();

  @override
  TransferConfirmationPage get transferConfirmationPage =>
      TransferConfirmationPageSw();

  @override
  Welcome get welcome => WelcomeSw();

  @override
  String get done => "Tayari";

  @override
  String isRequired(String name) => "$name: Ni Lazima Ujaze";

  @override
  String get home => "Nyumbani";

  @override
  String get account => "Akaunti";

  @override
  String get hide => "Ficha";

  @override
  String get next => "Endelea";

  @override
  String get no => "Hapana";

  @override
  String get refresh => "Onyesha upya";

  @override
  String get confirm => "Thibitisha";

  @override
  String get close => "Funga";

  @override
  String get unknownError => "Kuna tatizo limetokea";

  @override
  String get from => "Kutoka";

  @override
  String get phone => "Namba Ya Simu";

  @override
  String get yes => "Ndio";

  @override
  String get comingSoon => "Inakuja Hivi Karibuni!";

  @override
  String get show => "Onesha";

  @override
  String get beCareful => "Kuwa Makini!";

  @override
  String get seeAll => "Ona Yote";

  @override
  String get email => "Barua Pepe";

  @override
  String get send => "Tuma";

  @override
  String get remove => "Toa";

  @override
  String get to => "Kwenda";

  @override
  String get amount => "Kiasi";

  @override
  String get preview => "Angalia";

  @override
  String get tryAgain => "Jaribu Tena";

  @override
  String get cancel => "Ghairi";

  @override
  String get sendLink => "Tuma link kwa mteja";

  @override
  String get validationError => "Tafadhari angalia taarifa zako vizuri";
}

class RequestCardSw extends RequestCard {
  @override
  String get me => "Omba Kadi";

  @override
  String get desc =>
      "Pata kadi ya benki unayoweza kutumia kutoa pesa kutoka kwa ATM yoyote ya VISA, mahali popote na kufanya malipo mtandaoni.";

  @override
  String get action => "Tuma Ombi";
}

class GotYourRequestSw extends GotYourRequest {
  @override
  String get desc =>
      "Tumepokea ombi la kadi yako. Tutawasiliana nawe baada ya muda mfupi.";
}

class FundTopUpSw extends FundTopUp {
  @override
  String get expectedArrivalTime => "Muda itakapofika";

  @override
  String get fees => "Makato";

  @override
  String get telecom => "Mtandao";

  @override
  String get confirmWithdraw => "Thibitisha Muamala";

  @override
  String get transactionType => "Aina ya Muamala";

  @override
  String get title => "Ongeza Balance Yako";

  @override
  String get now => "Sasa Hivi";
}

class FundTopUpByPhoneSw extends FundTopUpByPhone {
  @override
  String get halopesaComingSoon => "HALOPESA inakuja hivi karibuni";

  @override
  String get productNameNotFilledError => "Tafadhari jaza jina la bidhaa";

  @override
  String get amountNotFilledError => "Hakikisha umejaza kiasi unachotaka utume";

  @override
  String get willEnablePayingThirdPartiesDesc =>
      "Tutakuwezesha kulipa kwenda LUKU, DSTV n.k. kupitia LipaLink hivi karibuni.";

  @override
  String get phoneNumber => "Namba ya simu";

  @override
  String get notifyCustomerDialogDesc =>
      "Mjulishe mteja kwamba ataombwa kuweka PIN kwa ajili ya kuthibitisha malipo";

  @override
  String get mpesaComingSoon => "MPESA inakuja hivi karibuni";

  @override
  String get title => "Kwa Namba ya Simu";

  @override
  String get azamPesaComingSoon => "AZAM-PESA inakuja hivi karibuni";

  @override
  String get unsupportedMomowalletError =>
      "Hakikisha namba yako ni ya TIGO au Airtel";

  @override
  String get tigoNumberHint => "Weka namba ya TIGO hapa";

  @override
  String get customerAskedForPindesc =>
      "Mteja ameombwa kuweka PIN ili kuthibitisha malipo. Refresh kuona kama malipo yamekamilika.";

  @override
  String get getPaymentsFrom => "Weka Pesa Kupitia:";

  @override
  String get notifyCustomerDialogTitle => "Mjulishe Mteja";

  @override
  String get airtelNumberHint => "Weka namba ya AIRTEL hapa";
}

class FundTopUpByBankSw extends FundTopUpByBank {
  @override
  String get title => "Kwa Benki";

  @override
  String get bankDetails => "Taarifa za Benki";

  @override
  String get bankName => "Jina la Benki";

  @override
  String get accountNumber => "Namba ya Akaunti";

  @override
  String get accountName => "Jina la Akaunti";

  @override
  String get desc =>
      "Unaweza kutuma pesa moja kwa moja kwenye akaunti yako ya benki kwa kutembelea benki, kutuma kutoka benki nyingine au kupitia mawakala wa EcoBank.";
}

class BalanceSw extends Balance {
  @override
  String get me => "Salio";

  @override
  String get failedToLoadError =>
      "Imeshindwa kupakia salio. Tafadhali jaribu tena baadaye.";

  @override
  String get insufficientFunds =>
      "Hela hazitoshi. Tafadhali jaza salio lako na ujaribu tena.";

  @override
  String get topUpByPhoneUssdmessage =>
      "Tumetuma USSD push ili kukamilisha ombi lako la kuongeza.";
}

class QuickActionsSw extends QuickActions {
  @override
  String get topUp => "Ongeza Salio";

  @override
  String get send => "Tuma Pesa";

  @override
  String get pay => "Lipa";
}

class HomeCardsSw extends HomeCards {
  @override
  String get title => "Actions";

  @override
  String get requestCardAction => "Omba Kadi";

  @override
  String get createGoalsAction => "Tengeneza Lengo";

  @override
  String get earnRewardsAction => "Pata Zawadi";
}

class ProfileSw extends Profile {
  @override
  String get accountNo => "Namba ya Akaunti";

  @override
  String get firstName => "Jina la Kwanza";

  @override
  String get lastName => "Jina la Ukoo";

  @override
  String get email => "Barua Pepe";

  @override
  String get businessName => "Jina la Biashara";

  @override
  String get name => "Jina";

  @override
  String get phone => "Namba Ya Simu";
}

class ProfilePageSw extends ProfilePage {
  @override
  String get changeBusinessName => "Badilisha Jina la Biashara";

  @override
  String notifyEmailBeingUsed(String email) => "You're logged in with $email";

  @override
  String get invalidPhoneNumberError =>
      "Tafadhari hakikisha umeweka namba ya simu sahihi";

  @override
  String get settings => "Mipangilio";

  @override
  String get changePindesc =>
      "Kubadilisha PIN kutakulazimu uwe signed-out kwenye application, na hivyo itakulazimu ku-sign in upya tena. Upo tayari kuendelea?";

  @override
  String get signOutDesc =>
      "Upo karibu kutoka katika application. Muda mwingine ukitaka kuingia katika application, utahitaji u-sign-in. Upo tayari kuendelea";

  @override
  String get createAccountPageTitle => "Jaza Taarifa Zako";

  @override
  String get signOut => "Sign Out";

  @override
  String optionalField(String field) => "$field (Hiari)";

  @override
  String get addPhone => "Weka Namba Ya Simu";

  @override
  String get changePin => "Badilisha PIN";

  @override
  String get changeLanguage => "Badilisha Lugha";

  @override
  String get businessNameBeingProfileNameError =>
      "Hakikisha jina la biashara sio tu muunganisho wa majina yako";
}

class SplashPageSw extends SplashPage {
  @override
  String get signInWithApple => "Ingia na Apple";

  @override
  String get profileFetchError => "Hatukuweza kupata wasifu wako";

  @override
  String get title => "Njia bora zaidi ya kusimamia pesa zako";

  @override
  String get signInWithGoogle => "Ingia na Google";
}

class TransactionsSw extends Transactions {
  @override
  String get addTransactionAction => "Anza Kutuma Pesa";

  @override
  String get me => "Miamala";

  @override
  String get loadingMsg => "Inapakia miamala...";

  @override
  String get failedToLoadError =>
      "Imeshindwa kupakia miamala. Tafadhali jaribu tena baadae.";

  @override
  String get noTransactions =>
      "Hujafanya miamala yoyote. Shughuli zako za hivi karibuni zitaonekana hapa.";
}

class TransactionSw extends Transaction {
  @override
  String get me => "Muamala";
}

class TransferTypeSw extends TransferType {
  @override
  String get temboWallet => "Akaunti Nyingine ya Tembo";

  @override
  String get phone => "Kwenda Mtandao wa Simu";

  @override
  String get bank => "Kwenda Benki";

  @override
  String get me => "Aina ya Muamala";
}

class TransferToBankPageSw extends TransferToBankPage {
  @override
  String get bankNotSelectedError => "Tafadhali hakikisha kuwa umechagua benki";

  @override
  String get selectBank => "Chagua Benki";

  @override
  String get amountError =>
      "Tafadhali hakikisha kuwa kiasi kimejazwa na ni sahihi.";
}

class TransferConfirmationPageSw extends TransferConfirmationPage {
  @override
  String get transferDetails => "Maelezo ya Muamala";

  @override
  String get recipientName => "Jina la Mpokeaji";

  @override
  String get sentMessage => "Muamala umeanzishwa!";

  @override
  String get title => "Thibitisha Muamala";
}

class WelcomeSw extends Welcome {
  @override
  String get heading => "Tunafurahi Kukuhudumia";

  @override
  String get desc =>
      '''TemboPlus ni huduma ya kifedha inayokuwezesha kufanya malipo kwenda China kwa usalama. 

Huduma hii inatolewa na kampuni yenye leseni kutoka Benki Kuu na unaweza kufanya malipo ya biashara yako moja kwa moja kwenda Akaunti za Benki zilizo China, Wechat au Alipay.''';

  @override
  String get registerAndroid =>
      "Hatua Zinazohitajika: Ingia kwa kutumia akaunti yako ya Google na Uthibitishe Nambari yako ya NIDA";

  @override
  String get registerIos =>
      "Hatua Zinazohitajika: Ingia kwa kutumia akaunti yako ya Google au iCloud na Uthibitishe Nambari yako ya NIDA";

  @override
  String greeting(String name) => "Habari $name";

  @override
  String get title => "Karibu";

  @override
  String get start =>
      "Bofya kitufe kilicho hapa chini ili kuanza kutumia huduma hii.";
}
