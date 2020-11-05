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

Future<HttpRequest> createStock(Key categoryKey, StockItem obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/${categoryKey.toJson()}";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateStock(Key categoryKey, Key key, StockItem obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/${categoryKey.toJson()}/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteStock(Key categoryKey, Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/${categoryKey.toJson()}/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
