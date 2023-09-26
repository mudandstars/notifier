import 'package:flutter/material.dart';
import 'package:notifier/type/webhook.dart';
import 'package:notifier/Components/webhooks_list.dart';
import 'package:notifier/Components/create_webhook_form.dart';

class Webhooks extends StatelessWidget {
  final List<Webhook>? webhooks;

  Webhooks({required this.webhooks});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Padding(
          padding: EdgeInsets.only(left: 20, right: 20, top: 15),
          child: CreateWebhookForm(),
        ),
        SizedBox(height: 50),
        WebhooksList(webhooks: webhooks),
      ],
    );
  }
}
