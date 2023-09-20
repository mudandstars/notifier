import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:namer_app/type/webhook.dart';

class ApiService {
  Future<List<Webhook>?> getWebhooks() async {
    var url = Uri.parse("http://127.0.0.1:6000/webhooks");
    var response = await http.get(url);

    if (response.statusCode == 200) {
      List<dynamic>? jsonList = jsonDecode(response.body)['webhooks'];
      print(jsonList);

      List<Webhook>? webhooks = jsonList
          ?.map((item) => Webhook(name: item["name"], url: item["url"]))
          .toList();
      return webhooks;
    }

    return null;
  }

  Future<bool> storeWebhook(String name) async {
    var url = Uri.parse("http://127.0.0.1:6000/webhooks");
    final headers = {
      'Content-Type': 'application/json',
    };
    final Map<String, dynamic> requestBody = {
      'Name': name,
    };

    var response =
        await http.post(url, headers: headers, body: jsonEncode(requestBody));

    if (response.statusCode == 200) {
      print("succes");
      return true;
    }

    return false;
  }
}
