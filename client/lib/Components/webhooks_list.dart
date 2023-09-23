import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:namer_app/API/api_service.dart';
import 'package:namer_app/State/global_state.dart';
import 'package:flutter/services.dart';
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
                        subtitle: Row(
                          children: [
                            Expanded(
                              child: Text(webhooks![index].url),
                            ),
                            IconButton(
                              icon: Icon(Icons.content_copy),
                              onPressed: () {
                                Clipboard.setData(ClipboardData(
                                        text: webhooks![index].url))
                                    .then((result) {
                                  ScaffoldMessenger.of(context).showSnackBar(
                                    SnackBar(
                                        content: Text('Copied to clipboard')),
                                  );
                                }).catchError((error) {
                                  ScaffoldMessenger.of(context).showSnackBar(
                                    SnackBar(content: Text('Copy failed')),
                                  );
                                });
                              },
                            ),
                          ],
                        ),
                        trailing: ElevatedButton(
                          onPressed: () async {
                            int id = webhooks![index].id;
                            bool isDeleted =
                                await ApiService().deleteWebhook(id);
                            print(
                                'Delete button pressed with $id, it was $isDeleted');

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
