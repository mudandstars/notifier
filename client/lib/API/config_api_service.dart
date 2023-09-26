import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:notifier/type/config.dart';
import 'package:notifier/type/webhook.dart';

class ConfigApiService {
  Future<Config?> get() async {
    var url = Uri.parse("http://127.0.0.1:6000/config");
    var response = await http.get(url);

    // if (response.statusCode == 200) {
    //   List<dynamic>? jsonList = jsonDecode(response.body)['webhooks'];

    //   List<Webhook>? webhooks = jsonList
    //       ?.map((item) =>
    //           Webhook(name: item["name"], url: item["url"], id: item["id"]))
    //       .toList();
    //   return webhooks;
    // }

    return null;
  }

  Future<bool> store(Config config) async {
    var url = Uri.parse("http://127.0.0.1:6000/config");
    final headers = {
      'Content-Type': 'application/json',
    };
    final Map<String, dynamic> requestBody = {
      'ngrokAuthToken': config.ngrokAuthToken,
      'ngrokPublicUrl': config.ngrokPublicUrl,
    };

    var response =
        await http.post(url, headers: headers, body: jsonEncode(requestBody));

    if (response.statusCode == 200) {
      return true;
    }

    return false;
  }

  Future<bool> delete(int id) async {
    var url = Uri.parse("http://127.0.0.1:6000/config");
    var response = await http.delete(url);

    if (response.statusCode == 200) {
      return true;
    }

    return false;
  }
}
