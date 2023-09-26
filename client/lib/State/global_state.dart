import 'package:notifier/API/config_api_service.dart';
import 'package:notifier/type/config.dart';
import 'package:notifier/type/webhook.dart';
import '../API/webhook_api_service.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<Webhook>? webhooks;
  Config? config;
  bool queriedBackend = false;

  void initState() async {
    webhooks = await WebhookApiService().get();
    config = await ConfigApiService().get();

    queriedBackend = true;

    notifyListeners();
  }
}
