import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/service.dart';

Future<HttpRequest> createService(Service obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/services";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateService(Key key, Service obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/services/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteService(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/services/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
