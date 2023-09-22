import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:namer_app/API/api_service.dart';
import 'package:namer_app/State/global_state.dart';
import 'package:namer_app/type/webhook.dart';

class WebhooksList extends StatelessWidget {
  final List<Webhook>? webhooks;

  WebhooksList({required this.webhooks});

  @override
  Widget build(BuildContext context) {
    var appState = context.watch<GlobalState>();

    return webhooks == null
        ? Text("There are no webhooks yet...")
        : webhooks!.isEmpty
            ? Center(
                child: CircularProgressIndicator(),
              )
            : Expanded(
                child: ListView.builder(
                  itemCount: webhooks!.length,
                  itemBuilder: (context, index) {
                    return Card(
                      child: ListTile(
                        title: Text(webhooks![index].name),
                        subtitle: Text(webhooks![index].url),
                        trailing: ElevatedButton(
                          onPressed: () async {
                            int id = webhooks![index].id;
                            bool isDeleted =
                                await ApiService().deleteWebhook(id);
                            print('Delete button pressed with $id, it was $isDeleted');

                            if (isDeleted) {
                              appState.initState();
                            }
                          },
                          child: Text('Delete'),
                        ),
                      ),
                    );
                  },
                ),
              );
  }
}
