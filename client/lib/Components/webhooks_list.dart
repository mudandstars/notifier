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
                    return Padding(
                        padding: EdgeInsets.only(left: 20, right: 20),
                        child: Card(
                          child: ListTile(
                            title: Text(webhooks![index].name),
                            subtitle: Row(
                              crossAxisAlignment: CrossAxisAlignment.center,
                              children: [
                                SizedBox(
                                  width: 350,
                                  child: Text(
                                    webhooks![index].url,
                                    overflow: TextOverflow.ellipsis,
                                    maxLines: 1,
                                  ),
                                ),
                                SizedBox(
                                  width: 10,
                                ),
                                SizedBox(
                                    width: 30,
                                    height: 30,
                                    child: IconButton(
                                      iconSize: 15,
                                      icon: Icon(Icons.content_copy),
                                      onPressed: () {
                                        Clipboard.setData(ClipboardData(
                                                text: webhooks![index].url))
                                            .then((result) {
                                          String webhookName =
                                              webhooks![index].name;
                                          ScaffoldMessenger.of(context)
                                              .showSnackBar(
                                            SnackBar(
                                                content: Text(
                                                    "Copied $webhookName's url to your clipboard.")),
                                          );
                                        }).catchError((error) {
                                          ScaffoldMessenger.of(context)
                                              .showSnackBar(
                                            SnackBar(
                                                content: Text('Copy failed')),
                                          );
                                        });
                                      },
                                    )),
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
                        ));
                  },
                ),
              );
  }
}
