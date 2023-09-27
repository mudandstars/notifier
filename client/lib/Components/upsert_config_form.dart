import 'package:flutter/material.dart';
import 'package:notifier/API/config_api_service.dart';
import 'package:notifier/Components/standard_button.dart';
import 'package:notifier/State/global_state.dart';
import 'package:notifier/type/config.dart';
import 'package:notifier/utils/notify.dart';
import 'package:provider/provider.dart';

class UpsertConfigForm extends StatelessWidget {
  final TextEditingController authTokenController = TextEditingController();
  final TextEditingController publicUrlController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    authTokenController.text = appState.config?.ngrokAuthToken ?? "";
    publicUrlController.text = appState.config?.ngrokPublicUrl ?? "";

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
            bool isUpserted = await ConfigApiService().upsert(Config(
                ngrokAuthToken: authTokenController.text,
                ngrokPublicUrl: publicUrlController.text));
            if (isUpserted) {
              appState.initState();
              appState.getConfig();
              notify(context,
                  "Saved config. It will spin up the new tunnel within 5 seconds. If it does not work, restart the backend.");
            } else {
              notify(context, "Failed to save config");
            }
          },
          text: 'Save',
        ),
      ],
    ));
  }
}
