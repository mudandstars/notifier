import 'dart:convert';

import 'package:http/http.dart' as http;

class ApiService {
  Future<List<String>?> getWebhooks() async {
    var url = Uri.parse("http://127.0.0.1:6000/webhooks");
    var response = await http.get(url);

    if (response.statusCode == 200) {
      List<dynamic> jsonList = jsonDecode(response.body)['webhooks'];
      List<String> webhooks = jsonList.map((item) => item.toString()).toList();
      return webhooks;
    }

    return null;
  }
}
