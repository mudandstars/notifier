import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:notifier/type/config.dart';

class ConfigApiService {
  Future<Config?> get() async {
    var url = Uri.parse("http://127.0.0.1:6000/config");
    var response = await http.get(url);

    if (response.statusCode == 200) {
      dynamic decodedJson = jsonDecode(response.body);

      return Config(
          ngrokAuthToken: decodedJson["NgrokAuthToken"],
          ngrokPublicUrl: decodedJson["NgrokPublicUrl"]);
    }

    return null;
  }

  Future<bool> upsert(Config config) async {
    var url = Uri.parse("http://127.0.0.1:6000/config");
    final headers = {
      'Content-Type': 'application/json',
    };
    final Map<String, dynamic> requestBody = {
      'ngrokAuthToken': config.ngrokAuthToken,
      'ngrokPublicUrl': config.ngrokPublicUrl,
    };

    var response =
        await http.put(url, headers: headers, body: jsonEncode(requestBody));

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
