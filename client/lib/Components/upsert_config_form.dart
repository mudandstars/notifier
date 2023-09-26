import 'package:flutter/material.dart';
import 'package:notifier/API/webhook_api_service.dart';
import 'package:notifier/Components/standard_button.dart';
import 'package:notifier/State/global_state.dart';
import 'package:provider/provider.dart';

class UpsertConfigForm extends StatelessWidget {
  final TextEditingController authTokenController = TextEditingController();
  final TextEditingController publicUrlController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    return SizedBox(
        child: Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Update Config",
          style: TextStyle(fontWeight: FontWeight.bold, fontSize: 22),
        ),
        SizedBox(
          height: 10,
        ),
        Row(mainAxisAlignment: MainAxisAlignment.end, children: [
          Expanded(
              child: TextField(
            controller: authTokenController,
            decoration: InputDecoration(
              border: UnderlineInputBorder(),
              hintText: "Auth token",
            ),
          )),
        ]),
        Row(mainAxisAlignment: MainAxisAlignment.end, children: [
          Expanded(
              child: TextField(
            controller: publicUrlController,
            decoration: InputDecoration(
              border: UnderlineInputBorder(),
              hintText: "Public Url",
            ),
          )),
        ]),
        StandardButton(
          onPressed: () async {
            // bool isStored = await ApiService().storeWebhook(authTokenController.text, publicUrlController.text);

            // if (isStored) {
            //   appState.initState();
            // }
          },
          text: 'Save',
        ),
      ],
    ));
  }
}
