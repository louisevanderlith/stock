import 'dart:convert';
import 'dart:html';

import 'package:mango_stock/bodies/category.dart';
import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

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

Future<HttpRequest> createStock(String category, StockItem obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/${category}";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateStock(String category, Key key, StockItem obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/${category}/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteStock(String category, Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/${category}/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
