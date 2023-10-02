import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'State/global_state.dart';
import 'Pages/home_page.dart';

void main() {
  runApp(NotifierApp());
}

class NotifierApp extends StatelessWidget {
  const NotifierApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => GlobalState(),
      child: MaterialApp(
        title: 'Notifier',
        theme: ThemeData(
          useMaterial3: true,
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.blueGrey),
        ),
        home: HomePage(),
      ),
    );
  }
}
