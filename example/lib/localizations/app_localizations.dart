// ignore_for_file: depend_on_referenced_packages

import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:intl/intl.dart' as intl;

import 'app_localizations_en.dart';
import 'app_localizations_sw.dart';

abstract class AppLocalizations {
  AppLocalizations(String locale)
      : localeName = intl.Intl.canonicalizedLocale(locale.toString());

  final String localeName;

  static AppLocalizations? of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }

  static const LocalizationsDelegate<AppLocalizations> delegate =
      _AppLocalizationsDelegate();

  static const List<LocalizationsDelegate<dynamic>> localizationsDelegates =
      <LocalizationsDelegate<dynamic>>[
    delegate,
    GlobalMaterialLocalizations.delegate,
    GlobalCupertinoLocalizations.delegate,
    GlobalWidgetsLocalizations.delegate,
  ];

  static const List<Locale> supportedLocales = <Locale>[
    Locale('en'),
    Locale('sw')
  ];

  RequestCard get requestCard;

  GotYourRequest get gotYourRequest;

  FundTopUp get fundTopUp;

  FundTopUpByPhone get fundTopUpByPhone;

  FundTopUpByBank get fundTopUpByBank;

  Balance get balance;

  QuickActions get quickActions;

  HomeCards get homeCards;

  Profile get profile;

  ProfilePage get profilePage;

  SplashPage get splashPage;

  Transactions get transactions;

  Transaction get transaction;

  TransferType get transferType;

  TransferToBankPage get transferToBankPage;

  TransferConfirmationPage get transferConfirmationPage;

  Welcome get welcome;

  /// No description provided for @done
  ///
  /// In en it is translated to:
  /// **Done**
  String get done;

  /// Used in forms when a certain required field is not filled
  ///
  /// In en it is translated to:
  /// **{name} is required**
  String isRequired(String name);

  /// No description provided for @home
  ///
  /// In en it is translated to:
  /// **Home**
  String get home;

  /// No description provided for @account
  ///
  /// In en it is translated to:
  /// **Account**
  String get account;

  /// No description provided for @hide
  ///
  /// In en it is translated to:
  /// **Hide**
  String get hide;

  /// No description provided for @next
  ///
  /// In en it is translated to:
  /// **Next**
  String get next;

  /// No description provided for @no
  ///
  /// In en it is translated to:
  /// **No**
  String get no;

  /// No description provided for @refresh
  ///
  /// In en it is translated to:
  /// **Refresh**
  String get refresh;

  /// No description provided for @confirm
  ///
  /// In en it is translated to:
  /// **Confirm**
  String get confirm;

  /// No description provided for @close
  ///
  /// In en it is translated to:
  /// **Close**
  String get close;

  /// No description provided for @unknownError
  ///
  /// In en it is translated to:
  /// **An unknown error happened**
  String get unknownError;

  /// No description provided for @from
  ///
  /// In en it is translated to:
  /// **From**
  String get from;

  /// No description provided for @phone
  ///
  /// In en it is translated to:
  /// **Phone**
  String get phone;

  /// No description provided for @yes
  ///
  /// In en it is translated to:
  /// **Yes**
  String get yes;

  /// No description provided for @comingSoon
  ///
  /// In en it is translated to:
  /// **Coming Soon!**
  String get comingSoon;

  /// No description provided for @show
  ///
  /// In en it is translated to:
  /// **Show**
  String get show;

  /// No description provided for @beCareful
  ///
  /// In en it is translated to:
  /// **Be careful!**
  String get beCareful;

  /// No description provided for @seeAll
  ///
  /// In en it is translated to:
  /// **See All**
  String get seeAll;

  /// No description provided for @email
  ///
  /// In en it is translated to:
  /// **Email**
  String get email;

  /// No description provided for @send
  ///
  /// In en it is translated to:
  /// **Send**
  String get send;

  /// No description provided for @remove
  ///
  /// In en it is translated to:
  /// **Remove**
  String get remove;

  /// No description provided for @to
  ///
  /// In en it is translated to:
  /// **To**
  String get to;

  /// No description provided for @amount
  ///
  /// In en it is translated to:
  /// **Amount**
  String get amount;

  /// No description provided for @preview
  ///
  /// In en it is translated to:
  /// **Preview**
  String get preview;

  /// No description provided for @tryAgain
  ///
  /// In en it is translated to:
  /// **Try Again**
  String get tryAgain;

  /// No description provided for @cancel
  ///
  /// In en it is translated to:
  /// **Cancel**
  String get cancel;

  /// No description provided for @sendLink
  ///
  /// In en it is translated to:
  /// **Send link to the customer**
  String get sendLink;

  /// No description provided for @validationError
  ///
  /// In en it is translated to:
  /// **Please check your information carefully**
  String get validationError;
}

class _AppLocalizationsDelegate
    extends LocalizationsDelegate<AppLocalizations> {
  const _AppLocalizationsDelegate();

  @override
  Future<AppLocalizations> load(Locale locale) {
    return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
  }

  @override
  bool shouldReload(_AppLocalizationsDelegate old) => false;

  @override
  bool isSupported(Locale locale) =>
      <String>['en', 'sw'].contains(locale.languageCode);

  AppLocalizations lookupAppLocalizations(Locale locale) {
    // Lookup logic when only language code is specified.
    switch (locale.languageCode) {
      case 'en':
        return AppLocalizationsEn();
      case 'sw':
        return AppLocalizationsSw();
    }

    throw FlutterError(
        'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
        'an issue with the localizations generation tool. Please file an issue '
        'on GitHub with a reproducible sample app and the gen-l10n configuration '
        'that was used.');
  }
}

abstract class RequestCard {
  /// No description provided for @me
  ///
  /// In en it is translated to:
  /// **Request Card**
  String get me;

  /// No description provided for @desc
  ///
  /// In en it is translated to:
  /// **Get a debit card that you can use to withdraw money from any VISA ATM, anywhere, and to make online payments.**
  String get desc;

  /// No description provided for @action
  ///
  /// In en it is translated to:
  /// **Request**
  String get action;
}

abstract class GotYourRequest {
  /// No description provided for @desc
  ///
  /// In en it is translated to:
  /// **We have received your card request. We will contact you shortly.**
  String get desc;
}

abstract class FundTopUp {
  /// No description provided for @expectedArrivalTime
  ///
  /// In en it is translated to:
  /// **Expected Arrival Time**
  String get expectedArrivalTime;

  /// No description provided for @fees
  ///
  /// In en it is translated to:
  /// **Fees**
  String get fees;

  /// No description provided for @telecom
  ///
  /// In en it is translated to:
  /// **Telecom Company**
  String get telecom;

  /// No description provided for @confirmWithdraw
  ///
  /// In en it is translated to:
  /// **Confirm Transaction**
  String get confirmWithdraw;

  /// No description provided for @transactionType
  ///
  /// In en it is translated to:
  /// **Transaction Type**
  String get transactionType;

  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **Top up your balance**
  String get title;

  /// No description provided for @now
  ///
  /// In en it is translated to:
  /// **Now**
  String get now;
}

abstract class FundTopUpByPhone {
  /// No description provided for @halopesaComingSoon
  ///
  /// In en it is translated to:
  /// **HALOPESA is coming soon**
  String get halopesaComingSoon;

  /// No description provided for @productNameNotFilledError
  ///
  /// In en it is translated to:
  /// **Please provide a provide name**
  String get productNameNotFilledError;

  /// No description provided for @amountNotFilledError
  ///
  /// In en it is translated to:
  /// **Make sure you fill in the amount you want to send**
  String get amountNotFilledError;

  /// No description provided for @willEnablePayingThirdPartiesDesc
  ///
  /// In en it is translated to:
  /// **We will enable you to pay to LUKU, DSTV etc. via LipaLink soon.**
  String get willEnablePayingThirdPartiesDesc;

  /// No description provided for @phoneNumber
  ///
  /// In en it is translated to:
  /// **Phone Number**
  String get phoneNumber;

  /// No description provided for @notifyCustomerDialogDesc
  ///
  /// In en it is translated to:
  /// **Inform the customer that they will be asked to enter a PIN to confirm payment**
  String get notifyCustomerDialogDesc;

  /// No description provided for @mpesaComingSoon
  ///
  /// In en it is translated to:
  /// **MPESA is coming soon**
  String get mpesaComingSoon;

  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **By Phone**
  String get title;

  /// No description provided for @azamPesaComingSoon
  ///
  /// In en it is translated to:
  /// **AZAM-PESA Coming Soon**
  String get azamPesaComingSoon;

  /// No description provided for @unsupportedMomowalletError
  ///
  /// In en it is translated to:
  /// **Make sure your phone number is a TIGO or AIRTEL number**
  String get unsupportedMomowalletError;

  /// No description provided for @tigoNumberHint
  ///
  /// In en it is translated to:
  /// **Fill your TIGO Number here**
  String get tigoNumberHint;

  /// No description provided for @customerAskedForPindesc
  ///
  /// In en it is translated to:
  /// **The customer is asked to enter a PIN to confirm the payment. Refresh to see if the payment is complete.**
  String get customerAskedForPindesc;

  /// No description provided for @getPaymentsFrom
  ///
  /// In en it is translated to:
  /// **Add Money From:**
  String get getPaymentsFrom;

  /// No description provided for @notifyCustomerDialogTitle
  ///
  /// In en it is translated to:
  /// **Notify The Customer**
  String get notifyCustomerDialogTitle;

  /// No description provided for @airtelNumberHint
  ///
  /// In en it is translated to:
  /// **Fill your AIRTEL here**
  String get airtelNumberHint;
}

abstract class FundTopUpByBank {
  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **By Bank**
  String get title;

  /// No description provided for @bankDetails
  ///
  /// In en it is translated to:
  /// **Bank Details**
  String get bankDetails;

  /// No description provided for @bankName
  ///
  /// In en it is translated to:
  /// **Bank Name**
  String get bankName;

  /// No description provided for @accountNumber
  ///
  /// In en it is translated to:
  /// **Account Number**
  String get accountNumber;

  /// No description provided for @accountName
  ///
  /// In en it is translated to:
  /// **Account Name**
  String get accountName;

  /// No description provided for @desc
  ///
  /// In en it is translated to:
  /// **You can send money directly to your bank account either by visiting the bank, sending from other banks or through EcoBank agents.**
  String get desc;
}

abstract class Balance {
  /// No description provided for @me
  ///
  /// In en it is translated to:
  /// **Balance**
  String get me;

  /// No description provided for @failedToLoadError
  ///
  /// In en it is translated to:
  /// **Failed to load balance. Please try again later.**
  String get failedToLoadError;

  /// No description provided for @insufficientFunds
  ///
  /// In en it is translated to:
  /// **Insufficient funds. Please top up your balance and try again.**
  String get insufficientFunds;

  /// No description provided for @topUpByPhoneUssdmessage
  ///
  /// In en it is translated to:
  /// **We have sent USSD push to complete your top up request.**
  String get topUpByPhoneUssdmessage;
}

abstract class QuickActions {
  /// No description provided for @topUp
  ///
  /// In en it is translated to:
  /// **Top Up**
  String get topUp;

  /// No description provided for @send
  ///
  /// In en it is translated to:
  /// **Send Money**
  String get send;

  /// No description provided for @pay
  ///
  /// In en it is translated to:
  /// **Pay**
  String get pay;
}

abstract class HomeCards {
  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **Actions**
  String get title;

  /// No description provided for @requestCardAction
  ///
  /// In en it is translated to:
  /// **Request Card**
  String get requestCardAction;

  /// No description provided for @createGoalsAction
  ///
  /// In en it is translated to:
  /// **Create Goals**
  String get createGoalsAction;

  /// No description provided for @earnRewardsAction
  ///
  /// In en it is translated to:
  /// **Earn Rewards**
  String get earnRewardsAction;
}

abstract class Profile {
  /// No description provided for @accountNo
  ///
  /// In en it is translated to:
  /// **Account Number**
  String get accountNo;

  /// No description provided for @firstName
  ///
  /// In en it is translated to:
  /// **First Name**
  String get firstName;

  /// No description provided for @lastName
  ///
  /// In en it is translated to:
  /// **Last Name**
  String get lastName;

  /// No description provided for @email
  ///
  /// In en it is translated to:
  /// **Email**
  String get email;

  /// No description provided for @businessName
  ///
  /// In en it is translated to:
  /// **Business Name**
  String get businessName;

  /// No description provided for @name
  ///
  /// In en it is translated to:
  /// **Name**
  String get name;

  /// No description provided for @phone
  ///
  /// In en it is translated to:
  /// **Phone Number**
  String get phone;
}

abstract class ProfilePage {
  /// No description provided for @changeBusinessName
  ///
  /// In en it is translated to:
  /// **Change Business Name**
  String get changeBusinessName;

  /// No description provided for @notifyEmailBeingUsed
  ///
  /// In en it is translated to:
  /// **You're logged in with {email}**
  String notifyEmailBeingUsed(String email);

  /// No description provided for @invalidPhoneNumberError
  ///
  /// In en it is translated to:
  /// **Please make sure you have entered the correct phone number**
  String get invalidPhoneNumberError;

  /// No description provided for @settings
  ///
  /// In en it is translated to:
  /// **Settings**
  String get settings;

  /// No description provided for @changePindesc
  ///
  /// In en it is translated to:
  /// **Changing the PIN will require you to be signed-out of the application, and thus you will have to sign in again. Are you ready to continue?**
  String get changePindesc;

  /// No description provided for @signOutDesc
  ///
  /// In en it is translated to:
  /// **You are about to exit the application. The next time you want to enter the application, you will need to sign-in. Are you ready to continue?**
  String get signOutDesc;

  /// No description provided for @createAccountPageTitle
  ///
  /// In en it is translated to:
  /// **Fill your details**
  String get createAccountPageTitle;

  /// No description provided for @signOut
  ///
  /// In en it is translated to:
  /// **Sign Out**
  String get signOut;

  /// No description provided for @optionalField
  ///
  /// In en it is translated to:
  /// **{field} (Optional)**
  String optionalField(String field);

  /// No description provided for @addPhone
  ///
  /// In en it is translated to:
  /// **Add Phone Number**
  String get addPhone;

  /// No description provided for @changePin
  ///
  /// In en it is translated to:
  /// **Change PIN**
  String get changePin;

  /// No description provided for @changeLanguage
  ///
  /// In en it is translated to:
  /// **Change Language**
  String get changeLanguage;

  /// No description provided for @businessNameBeingProfileNameError
  ///
  /// In en it is translated to:
  /// **Please make sure a business name isn't just a combination of your first and last names**
  String get businessNameBeingProfileNameError;
}

abstract class SplashPage {
  /// No description provided for @signInWithApple
  ///
  /// In en it is translated to:
  /// **Sign In With Apple**
  String get signInWithApple;

  /// No description provided for @profileFetchError
  ///
  /// In en it is translated to:
  /// **We could not fetch your profile**
  String get profileFetchError;

  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **A better way to manage your money**
  String get title;

  /// No description provided for @signInWithGoogle
  ///
  /// In en it is translated to:
  /// **Sign In With Google**
  String get signInWithGoogle;
}

abstract class Transactions {
  /// No description provided for @addTransactionAction
  ///
  /// In en it is translated to:
  /// **Send Sending Money**
  String get addTransactionAction;

  /// No description provided for @me
  ///
  /// In en it is translated to:
  /// **Transactions**
  String get me;

  /// No description provided for @loadingMsg
  ///
  /// In en it is translated to:
  /// **Loading transactions...**
  String get loadingMsg;

  /// No description provided for @failedToLoadError
  ///
  /// In en it is translated to:
  /// **Failed to load transactions. Please try again later.**
  String get failedToLoadError;

  /// No description provided for @noTransactions
  ///
  /// In en it is translated to:
  /// **You have not done any transactions. Your recent activities will appear here.**
  String get noTransactions;
}

abstract class Transaction {
  /// No description provided for @me
  ///
  /// In en it is translated to:
  /// **Transaction**
  String get me;
}

abstract class TransferType {
  /// No description provided for @temboWallet
  ///
  /// In en it is translated to:
  /// **Tembo Wallet Transfer**
  String get temboWallet;

  /// No description provided for @phone
  ///
  /// In en it is translated to:
  /// **Mobile Wallet Transfer**
  String get phone;

  /// No description provided for @bank
  ///
  /// In en it is translated to:
  /// **Bank Transfer**
  String get bank;

  /// No description provided for @me
  ///
  /// In en it is translated to:
  /// **Transfer Type**
  String get me;
}

abstract class TransferToBankPage {
  /// No description provided for @bankNotSelectedError
  ///
  /// In en it is translated to:
  /// **Please make sure the bank is selected**
  String get bankNotSelectedError;

  /// No description provided for @selectBank
  ///
  /// In en it is translated to:
  /// **Select Bank**
  String get selectBank;

  /// No description provided for @amountError
  ///
  /// In en it is translated to:
  /// **Please make sure the amount is filled and valid.**
  String get amountError;
}

abstract class TransferConfirmationPage {
  /// No description provided for @transferDetails
  ///
  /// In en it is translated to:
  /// **Transfer Details**
  String get transferDetails;

  /// No description provided for @recipientName
  ///
  /// In en it is translated to:
  /// **Recipient's Name**
  String get recipientName;

  /// No description provided for @sentMessage
  ///
  /// In en it is translated to:
  /// **Transfer initiated!**
  String get sentMessage;

  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **Confirm Transfer**
  String get title;
}

abstract class Welcome {
  /// No description provided for @heading
  ///
  /// In en it is translated to:
  /// **We are happy to serve you**
  String get heading;

  /// No description provided for @desc
  ///
  /// In en it is translated to:
  /// **TemboPlus is a financial service that allows you to make payments to China safely. This service is provided by a company licensed by the Central Bank and you can make your business payments directly to Bank Accounts in China, Wechat or Alipay.**
  String get desc;

  /// No description provided for @registerAndroid
  ///
  /// In en it is translated to:
  /// **Required Steps: Log in with your Google Account & Verify Your NIN Number**
  String get registerAndroid;

  /// No description provided for @registerIos
  ///
  /// In en it is translated to:
  /// **Required Steps: Log in with your Google or iCloud Account & Verify Your NIN Number**
  String get registerIos;

  /// No description provided for @greeting
  ///
  /// In en it is translated to:
  /// **Hello {name}**
  String greeting(String name);

  /// No description provided for @title
  ///
  /// In en it is translated to:
  /// **Welcome**
  String get title;

  /// No description provided for @start
  ///
  /// In en it is translated to:
  /// **Anza kwa kuthibitisha namba yako ya simu**
  String get start;
}
