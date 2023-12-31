import 'package:flutter/material.dart';
import 'package:notifier/API/webhook_api_service.dart';
import 'package:notifier/Components/standard_button.dart';
import 'package:notifier/Components/text_input.dart';
import 'package:notifier/State/global_state.dart';
import 'package:notifier/utils/notify.dart';
import 'package:provider/provider.dart';

class CreateWebhookForm extends StatelessWidget {
  final TextEditingController textEditingController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    return SizedBox(
        child: Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(mainAxisAlignment: MainAxisAlignment.end, children: [
          Expanded(
              child: TextInput(
            controller: textEditingController,
            placeholder: "New Webhook Name..",
          )),
          SizedBox(
            width: 20,
          ),
          Container(
            height: 48,
            child: StandardButton(
              onPressed: () async {
                String name = textEditingController.text;
                bool isStored = await WebhookApiService().store(name);

                if (isStored) {
                  appState.initState();
                  notify(context, "Successfully created webhook");
                } else {
                  notify(context, "Failed creating webhook");
                }
              },
              text: 'Create',
            ),
          ),
        ])
      ],
    ));
  }
}
