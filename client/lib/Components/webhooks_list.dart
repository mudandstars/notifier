import 'package:flutter/material.dart';
import 'package:notifier/utils/store_to_clipboard.dart';
import 'package:provider/provider.dart';
import 'package:notifier/API/webhook_api_service.dart';
import 'package:notifier/State/global_state.dart';
import 'package:notifier/type/webhook.dart';

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
                        padding: EdgeInsets.only(left: 15, right: 15),
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
                                        storeToClipboard(
                                            context, webhooks![index]);
                                      },
                                    )),
                              ],
                            ),
                            trailing: SizedBox(
                                width: 40,
                                height: 40,
                                child: IconButton(
                                  iconSize: 20,
                                  icon: Icon(Icons.delete),
                                  onPressed: () async {
                                    int id = webhooks![index].id;
                                    bool isDeleted =
                                        await WebhookApiService().delete(id);

                                    if (isDeleted) {
                                      appState.initState();
                                    }
                                  },
                                )),
                          ),
                        ));
                  },
                ),
              );
  }
}
