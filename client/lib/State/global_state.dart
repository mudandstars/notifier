import '../API/ApiService.dart';
import 'package:flutter/material.dart';

class GlobalState extends ChangeNotifier {
  List<String> webhooks = [];

  void initState() async {
    webhooks = await getData();
    print(webhooks);

    notifyListeners();
  }

  Future<List<String>> getData() async {
    return await ApiService().getWebhooks() ?? [];
  }
}
