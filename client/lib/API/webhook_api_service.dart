import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:notifier/type/webhook.dart';

class WebhookApiService {
  Future<List<Webhook>?> getWebhooks() async {
    var url = Uri.parse("http://127.0.0.1:6000/webhooks");
    var response = await http.get(url);

    if (response.statusCode == 200) {
      List<dynamic>? jsonList = jsonDecode(response.body)['webhooks'];

      List<Webhook>? webhooks = jsonList
          ?.map((item) =>
              Webhook(name: item["name"], url: item["url"], id: item["id"]))
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
      return true;
    }

    return false;
  }

  Future<bool> deleteWebhook(int id) async {
    var url = Uri.parse("http://127.0.0.1:6000/webhooks/$id");
    var response = await http.delete(url);

    if (response.statusCode == 200) {
      return true;
    }

    return false;
  }
}
