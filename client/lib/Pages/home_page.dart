import 'package:flutter/material.dart';
import 'package:notifier/Pages/config.dart';
import 'package:notifier/Pages/webhooks.dart';
import 'package:provider/provider.dart';
import '../State/global_state.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final TextEditingController textEditingController = TextEditingController();
  int _selectedTabIndex = 0;

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    if (!appState.queriedWebhooks) {
      appState.initState();
    }

    return DefaultTabController(
      length: 2,
      initialIndex: _selectedTabIndex,
      child: Scaffold(
        appBar: AppBar(
          toolbarHeight: 3,
          bottom: TabBar(
            tabs: [
              Tab(text: 'Webhooks'),
              Tab(text: 'Config'),
            ],
            onTap: (index) {
              setState(() {
                _selectedTabIndex = index;
              });
            },
          ),
        ),
        body: TabBarView(
          children: [
            Webhooks(
              webhooks: appState.webhooks,
            ),
            Config(),
          ],
        ),
      ),
    );
  }
}
