import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/part.dart';

Future<HttpRequest> createPart(Part obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/parts";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updatePart(Key key, Part obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/parts/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deletePart(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/parts/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
