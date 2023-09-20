import 'package:flutter/material.dart';
import 'package:namer_app/Components/webhooks_list.dart';
import 'package:provider/provider.dart';
import '../State/global_state.dart';

class HomePage extends StatelessWidget {
  final TextEditingController textEditingController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    if (appState.webhooks.isEmpty) {
      appState.initState();
    }

    return Scaffold(
      key: Key(appState.webhooks.length.toString()),
      body: Column(
        children: [
          TextField(
            controller: textEditingController,
            decoration: InputDecoration(
              border: UnderlineInputBorder(),
              hintText: "Enter the project's name",
            ),
          ),
          Row(
            children: [
              ElevatedButton(
                onPressed: () {
                  // Access the current value of the text field
                  String currentValue = textEditingController.text;
                  print("Current Value: $currentValue");
                },
                child: Text('Get Value'),
              ),
              ElevatedButton(
                onPressed: () {
                  appState.initState();
                },
                child: Text('Refresh GET'),
              ),
            ],
          ),
          WebhooksList(webhooks: appState.webhooks)
        ],
      ),
    );
  }
}