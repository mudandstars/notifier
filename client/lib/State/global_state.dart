import 'package:notifier/type/webhook.dart';
import '../API/api_service.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<Webhook>? webhooks;
  bool queriedWebhooks = false;

  void initState() async {
    webhooks = await ApiService().getWebhooks();
    queriedWebhooks = true;

    notifyListeners();
  }
}
