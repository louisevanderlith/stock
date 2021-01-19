import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/category.dart';

Future<HttpRequest> searchCategories(Category obj) async {
  var apiroute = getEndpoint("stock");
  final jobj = jsonEncode(obj.toJson());
  final search = base64UrlEncode(utf8.encode(jobj));
  var url = "${apiroute}/categories/${search}";

  return invokeService("GET", url, null);
}

Future<HttpRequest> createCategory(Category obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateCategory(Key key, Category obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteCategory(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
