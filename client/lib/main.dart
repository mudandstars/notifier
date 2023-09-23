import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'State/global_state.dart';
import 'Pages/home_page.dart';
import 'dart:io';

import 'package:logging/logging.dart';

void main() {
  _configureLogging();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => GlobalState(),
      child: MaterialApp(
        title: 'Notifier',
        theme: ThemeData(
          useMaterial3: true,
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
        ),
        home: HomePage(),
      ),
    );
  }
}

void _configureLogging() {
  Logger.root.level = Level.ALL;

  final logFile = File('app_log.txt');

  Logger.root.onRecord.listen((record) {
    print('${record.level.name}: ${record.time}: ${record.message}');

    logFile.writeAsStringSync('${DateTime.now()}: $record\n',
        mode: FileMode.append);
  });
}
