import 'package:flutter/material.dart';
import 'package:namer_app/API/ApiService.dart';
import 'package:namer_app/State/global_state.dart';
import 'package:provider/provider.dart';

class CreateWebhookForm extends StatelessWidget {
  final TextEditingController textEditingController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    return SizedBox(
        height: 200.0,
        child: Row(children: [
          Expanded(
              child: TextField(
            controller: textEditingController,
            decoration: InputDecoration(
              border: UnderlineInputBorder(),
              hintText: "Enter the project's name",
            ),
          )),
          ElevatedButton(
            onPressed: () async {
              String name = textEditingController.text;
              bool isStored = await ApiService().storeWebhook(name);

              if (isStored) {
                appState.initState();
              }
            },
            child: Text('Create Webhook'),
          ),
        ]));
  }
}
