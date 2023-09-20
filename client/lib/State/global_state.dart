import '../API/ApiService.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<String>? webhooks;
  bool queriedWebhooks = false;

  void initState() async {
    webhooks = await getData();
    queriedWebhooks = true;
    print(webhooks);

    notifyListeners();
  }

  Future<List<String>?> getData() async {
    return await ApiService().getWebhooks();
  }
}
