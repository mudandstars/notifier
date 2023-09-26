import 'package:notifier/API/config_api_service.dart';
import 'package:notifier/type/config.dart';
import 'package:notifier/type/webhook.dart';
import '../API/webhook_api_service.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<Webhook>? webhooks;
  Config? config;
  bool queriedWebhooks = false;

  void initState() async {
    webhooks = await WebhookApiService().get();
    queriedWebhooks = true;

    notifyListeners();
  }

  void getConfig() async {
    config = await ConfigApiService().get();
  }
}
