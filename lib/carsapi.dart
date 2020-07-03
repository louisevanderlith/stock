import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/car.dart';

Future<HttpRequest> createCar(Car obj) async {
  var apiroute = getEndpoint("stock");
  var url = "${apiroute}/cars";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateCar(Key key, Car obj) async {
  var route = getEndpoint("stock");
  var url = "${route}/cars/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteCar(Key key) async {
  var route = getEndpoint("stock");
  var url = "${route}/cars/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
