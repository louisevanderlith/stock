import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/property.dart';

Future<HttpRequest> createProperty(Property obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/properties";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateProperty(Key key, Property obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/properties/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteProperty(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/properties/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
