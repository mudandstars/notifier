import 'package:flutter/material.dart';
import 'package:notifier/Components/webhooks_list.dart';
import 'package:notifier/Components/create_webhook_form.dart';
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
        key: Key(appState.webhooks?.length.toString() ?? "0"),
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
            Column(
              children: [
                Padding(
                  padding: EdgeInsets.only(left: 20, right: 20, top: 15),
                  child: CreateWebhookForm(),
                ),
                SizedBox(height: 50),
                WebhooksList(webhooks: appState.webhooks),
              ],
            ),
            Center(
              child: Text('Page 2 Content'),
            ),
          ],
        ),
      ),
    );
  }
}
