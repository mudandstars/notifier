import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:notifier/type/webhook.dart';

void storeToClipboard(BuildContext context, Webhook webhook) {
  Clipboard.setData(ClipboardData(text: webhook.url)).then((result) {
    String webhookName = webhook.name;
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text("Copied $webhookName's url to your clipboard.")),
    );
  }).catchError((error) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Copy failed')),
    );
  });
}
