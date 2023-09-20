import 'package:namer_app/type/webhook.dart';

import '../API/api_service.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<Webhook>? webhooks;
  bool queriedWebhooks = false;

  void initState() async {
    webhooks = await getData();
    queriedWebhooks = true;
    print(webhooks);

    notifyListeners();
  }

  Future<List<Webhook>?> getData() async {
    return await ApiService().getWebhooks();
  }
}
