import 'package:flutter/material.dart';
import 'package:namer_app/type/webhook.dart';

class WebhooksList extends StatelessWidget {
  final List<Webhook>? webhooks;

  WebhooksList({required this.webhooks});

  @override
  Widget build(BuildContext context) {
    return webhooks == null
        ? Text("there are no webhooks yet...")
        : webhooks!.isEmpty
            ? const Center(
                child: CircularProgressIndicator(),
              )
            : Expanded(
                child: SizedBox(
                    height: 200.0,
                    child: ListView.builder(
                      itemCount: webhooks!.length,
                      itemBuilder: (context, index) {
                        return ListTile(
                          title: Text(webhooks![index].name),
                          subtitle: Text(webhooks![index].url),
                        );
                      },
                    )),
              );
  }
}
