import 'app_localizations.dart';

/// The translations for en).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  RequestCard get requestCard => RequestCardEn();

  @override
  GotYourRequest get gotYourRequest => GotYourRequestEn();

  @override
  FundTopUp get fundTopUp => FundTopUpEn();

  @override
  FundTopUpByPhone get fundTopUpByPhone => FundTopUpByPhoneEn();

  @override
  FundTopUpByBank get fundTopUpByBank => FundTopUpByBankEn();

  @override
  Balance get balance => BalanceEn();

  @override
  QuickActions get quickActions => QuickActionsEn();

  @override
  HomeCards get homeCards => HomeCardsEn();

  @override
  Profile get profile => ProfileEn();

  @override
  ProfilePage get profilePage => ProfilePageEn();

  @override
  SplashPage get splashPage => SplashPageEn();

  @override
  Transactions get transactions => TransactionsEn();

  @override
  Transaction get transaction => TransactionEn();

  @override
  TransferType get transferType => TransferTypeEn();

  @override
  TransferToBankPage get transferToBankPage => TransferToBankPageEn();

  @override
  TransferConfirmationPage get transferConfirmationPage =>
      TransferConfirmationPageEn();

  @override
  Welcome get welcome => WelcomeEn();

  @override
  String get done => "Done";

  @override
  String isRequired(String name) => "$name is required";

  @override
  String get home => "Home";

  @override
  String get account => "Account";

  @override
  String get hide => "Hide";

  @override
  String get next => "Next";

  @override
  String get no => "No";

  @override
  String get refresh => "Refresh";

  @override
  String get confirm => "Confirm";

  @override
  String get close => "Close";

  @override
  String get unknownError => "An unknown error happened";

  @override
  String get from => "From";

  @override
  String get phone => "Phone";

  @override
  String get yes => "Yes";

  @override
  String get comingSoon => "Coming Soon!";

  @override
  String get show => "Show";

  @override
  String get beCareful => "Be careful!";

  @override
  String get seeAll => "See All";

  @override
  String get email => "Email";

  @override
  String get send => "Send";

  @override
  String get remove => "Remove";

  @override
  String get to => "To";

  @override
  String get amount => "Amount";

  @override
  String get preview => "Preview";

  @override
  String get tryAgain => "Try Again";

  @override
  String get cancel => "Cancel";

  @override
  String get sendLink => "Send link to the customer";

  @override
  String get validationError => "Please check your information carefully";
}

class RequestCardEn extends RequestCard {
  @override
  String get me => "Request Card";

  @override
  String get desc =>
      "Get a debit card that you can use to withdraw money from any VISA ATM, anywhere, and to make online payments.";

  @override
  String get action => "Request";
}

class GotYourRequestEn extends GotYourRequest {
  @override
  String get desc =>
      "We have received your card request. We will contact you shortly.";
}

class FundTopUpEn extends FundTopUp {
  @override
  String get expectedArrivalTime => "Expected Arrival Time";

  @override
  String get fees => "Fees";

  @override
  String get telecom => "Telecom Company";

  @override
  String get confirmWithdraw => "Confirm Transaction";

  @override
  String get transactionType => "Transaction Type";

  @override
  String get title => "Top up your balance";

  @override
  String get now => "Now";
}

class FundTopUpByPhoneEn extends FundTopUpByPhone {
  @override
  String get halopesaComingSoon => "HALOPESA is coming soon";

  @override
  String get productNameNotFilledError => "Please provide a provide name";

  @override
  String get amountNotFilledError =>
      "Make sure you fill in the amount you want to send";

  @override
  String get willEnablePayingThirdPartiesDesc =>
      "We will enable you to pay to LUKU, DSTV etc. via LipaLink soon.";

  @override
  String get phoneNumber => "Phone Number";

  @override
  String get notifyCustomerDialogDesc =>
      "Inform the customer that they will be asked to enter a PIN to confirm payment";

  @override
  String get mpesaComingSoon => "MPESA is coming soon";

  @override
  String get title => "By Phone";

  @override
  String get azamPesaComingSoon => "AZAM-PESA Coming Soon";

  @override
  String get unsupportedMomowalletError =>
      "Make sure your phone number is a TIGO or AIRTEL number";

  @override
  String get tigoNumberHint => "Fill your TIGO Number here";

  @override
  String get customerAskedForPindesc =>
      "The customer is asked to enter a PIN to confirm the payment. Refresh to see if the payment is complete.";

  @override
  String get getPaymentsFrom => "Add Money From:";

  @override
  String get notifyCustomerDialogTitle => "Notify The Customer";

  @override
  String get airtelNumberHint => "Fill your AIRTEL here";
}

class FundTopUpByBankEn extends FundTopUpByBank {
  @override
  String get title => "By Bank";

  @override
  String get bankDetails => "Bank Details";

  @override
  String get bankName => "Bank Name";

  @override
  String get accountNumber => "Account Number";

  @override
  String get accountName => "Account Name";

  @override
  String get desc =>
      "You can send money directly to your bank account either by visiting the bank, sending from other banks or through EcoBank agents.";
}

class BalanceEn extends Balance {
  @override
  String get me => "Balance";

  @override
  String get failedToLoadError =>
      "Failed to load balance. Please try again later.";

  @override
  String get insufficientFunds =>
      "Insufficient funds. Please top up your balance and try again.";

  @override
  String get topUpByPhoneUssdmessage =>
      "We have sent USSD push to complete your top up request.";
}

class QuickActionsEn extends QuickActions {
  @override
  String get topUp => "Top Up";

  @override
  String get send => "Send Money";

  @override
  String get pay => "Pay";
}

class HomeCardsEn extends HomeCards {
  @override
  String get title => "Actions";

  @override
  String get requestCardAction => "Request Card";

  @override
  String get createGoalsAction => "Create Goals";

  @override
  String get earnRewardsAction => "Earn Rewards";
}

class ProfileEn extends Profile {
  @override
  String get accountNo => "Account Number";

  @override
  String get firstName => "First Name";

  @override
  String get lastName => "Last Name";

  @override
  String get email => "Email";

  @override
  String get businessName => "Business Name";

  @override
  String get name => "Name";

  @override
  String get phone => "Phone Number";
}

class ProfilePageEn extends ProfilePage {
  @override
  String get changeBusinessName => "Change Business Name";

  @override
  String notifyEmailBeingUsed(String email) => "You're logged in with $email";

  @override
  String get invalidPhoneNumberError =>
      "Please make sure you have entered the correct phone number";

  @override
  String get settings => "Settings";

  @override
  String get changePindesc =>
      "Changing the PIN will require you to be signed-out of the application, and thus you will have to sign in again. Are you ready to continue?";

  @override
  String get signOutDesc =>
      "You are about to exit the application. The next time you want to enter the application, you will need to sign-in. Are you ready to continue?";

  @override
  String get createAccountPageTitle => "Fill your details";

  @override
  String get signOut => "Sign Out";

  @override
  String optionalField(String field) => "$field (Optional)";

  @override
  String get addPhone => "Add Phone Number";

  @override
  String get changePin => "Change PIN";

  @override
  String get changeLanguage => "Change Language";

  @override
  String get businessNameBeingProfileNameError =>
      "Please make sure a business name isn't just a combination of your first and last names";
}

class SplashPageEn extends SplashPage {
  @override
  String get signInWithApple => "Sign In With Apple";

  @override
  String get profileFetchError => "We could not fetch your profile";

  @override
  String get title => "A better way to manage your money";

  @override
  String get signInWithGoogle => "Sign In With Google";
}

class TransactionsEn extends Transactions {
  @override
  String get addTransactionAction => "Send Sending Money";

  @override
  String get me => "Transactions";

  @override
  String get loadingMsg => "Loading transactions...";

  @override
  String get failedToLoadError =>
      "Failed to load transactions. Please try again later.";

  @override
  String get noTransactions =>
      "You have not done any transactions. Your recent activities will appear here.";
}

class TransactionEn extends Transaction {
  @override
  String get me => "Transaction";
}

class TransferTypeEn extends TransferType {
  @override
  String get temboWallet => "Tembo Wallet Transfer";

  @override
  String get phone => "Mobile Wallet Transfer";

  @override
  String get bank => "Bank Transfer";

  @override
  String get me => "Transfer Type";
}

class TransferToBankPageEn extends TransferToBankPage {
  @override
  String get bankNotSelectedError => "Please make sure the bank is selected";

  @override
  String get selectBank => "Select Bank";

  @override
  String get amountError => "Please make sure the amount is filled and valid.";
}

class TransferConfirmationPageEn extends TransferConfirmationPage {
  @override
  String get transferDetails => "Transfer Details";

  @override
  String get recipientName => "Recipient's Name";

  @override
  String get sentMessage => "Transfer initiated!";

  @override
  String get title => "Confirm Transfer";
}

class WelcomeEn extends Welcome {
  @override
  String get heading => "We are happy to serve you";

  @override
  String get desc =>
      '''TemboPlus is a financial service that allows you to make payments to China safely. 

This service is provided by a company licensed by the Central Bank and you can make your business payments directly to Bank Accounts in China, Wechat or Alipay.''';

  @override
  String get registerAndroid =>
      "Required Steps: Log in with your Google Account & Verify Your NIN Number";

  @override
  String get registerIos =>
      "Required Steps: Log in with your Google or iCloud Account & Verify Your NIN Number";

  @override
  String greeting(String name) => "Hello $name";

  @override
  String get title => "Welcome";

  @override
  String get start => "Anza kwa kuthibitisha namba yako ya simu";
}
