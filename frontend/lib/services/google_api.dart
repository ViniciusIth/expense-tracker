import 'package:http/http.dart' as http;
import 'dart:convert';

import 'dart:async';

String sheetID = '1G0CA8kjCxX5ZYIp6TeHINEI__55uVARz-3DByz72nt0';
String deploymentID =
    'AKfycbyWKgMrrnSMtA4d2CJ3dh28NSVhHeeythzT0vqv1Zr0BEXuOD-9Jh5R0KT9N9avsw3x';

Future<Map> callWebAPP({required Map body}) async {
  Map dataDict = {};
  Uri URL = Uri.parse(
    'https://script.google.com/macros/s/${deploymentID}/exec',
  );

  try {
    await http
        .post(URL, body: body, encoding: Encoding.getByName('application/json'))
        .then((response) async {
          if ([200, 201].contains(response.statusCode)) {
            dataDict = jsonDecode(response.body);
          }

          if (response.statusCode == 302) {
            String redirectedUrl = response.headers['location'] ?? "";
            if (redirectedUrl.isNotEmpty) {
              Uri url = Uri.parse(redirectedUrl);
              await http.get(url).then((response) {
                if ([200, 201].contains(response.statusCode)) {
                  dataDict = jsonDecode(response.body);
                }
              });
            }
          } else {
            print("Other StatusCode: ${response.statusCode}");
          }
        });
  } catch (e) {
    print("Failed: $e");
  }

  return dataDict;
}

Future<Map> getSheetsData({required String action}) async {
  Map body = {"sheetID": sheetID, "action": action};

  Map dataDict = await callWebAPP(body: body);

  return dataDict;
}

Future<Map> appendToSheet({required List<dynamic> rowData}) async {
  Map body = {
    "sheetID": sheetID,
    "action": "append",
    "data": jsonEncode(rowData), // Encode the row data as a JSON string
  };

  Map dataDict = await callWebAPP(body: body);

  return dataDict;
}
