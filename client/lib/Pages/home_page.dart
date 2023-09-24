import 'package:flutter/material.dart';
import 'package:notifier/Components/webhooks_list.dart';
import 'package:notifier/Components/create_webhook_form.dart';
import 'package:provider/provider.dart';
import '../State/global_state.dart';

class HomePage extends StatelessWidget {
  final TextEditingController textEditingController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    if (!appState.queriedWebhooks) {
      appState.initState();
    }

    return Scaffold(
      key: Key(appState.webhooks?.length.toString() ?? "0"),
      body: Column(
        children: [
          Padding(
              padding: EdgeInsets.only(left: 20, right: 20),
              child: CreateWebhookForm()
              ),
          WebhooksList(webhooks: appState.webhooks),
        ],
      ),
    );
  }
}
