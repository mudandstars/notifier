import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:notifier/type/webhook.dart';
import 'package:notifier/utils/notify.dart';

void storeToClipboard(BuildContext context, Webhook webhook) {
  Clipboard.setData(ClipboardData(text: webhook.url)).then((result) {
    String webhookName = webhook.name;
    notify(context, "Copied $webhookName's url to your clipboard.");
  }).catchError((error) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Copy failed')),
    );
  });
}
