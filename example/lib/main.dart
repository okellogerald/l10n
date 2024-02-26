import 'package:flutter/material.dart';
import 'package:l10n_flutter/localizations/app_localizations.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      locale: const Locale("en"),
      supportedLocales: AppLocalizations.supportedLocales,
      localizationsDelegates: AppLocalizations.localizationsDelegates,
      home: Builder(builder: (context) {
        return Scaffold(
          body: Center(
            child: Text(AppLocalizations.of(context)!.home),
          ),
        );
      }),
    );
  }
}
