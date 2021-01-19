import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/product.dart';

Future<HttpRequest> searchProducts(Product obj) async {
  var apiroute = getEndpoint("stock");
  final jobj = jsonEncode(obj.toJson());
  final search = base64UrlEncode(utf8.encode(jobj));
  var url = "${apiroute}/products/${search}";

  return invokeService("GET", url, null);
}

Future<HttpRequest> createProduct(Product obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/products";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateProduct(Key key, Product obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/products/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteProduct(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/products/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
