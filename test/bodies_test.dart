import 'dart:convert';

import 'package:mango_stock/bodies/category.dart';
import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';
import 'package:test/test.dart';

void main() {
  group("Category", () {
    test("toJson single object", () {
      final items = new List<StockItem>();
      items.add(FakeItem());

      final obj = new Category("test", "some text", " the description",
          "http://some.where.com", 0, "www", new Key("0`0"), items);

      final actual = jsonEncode(obj);
      final match =
          '{"Name":"test","Text":"some text","Description":" the description","PageURL":"http://some.where.com","BaseCategory":0,"ClientID":"www","ImageKey":"0`0","Items":[{"ItemKey":"0`0","ImageKey":"0`0","OwnerKey":"0`0","Expires":"2020-06-06T00:00:00.000","Price":10,"Tags":["tag","tag"],"Location":"Somewhere","Views":0,"OwnerHistory":{},"EstimateValue":15,"Currency":"ZAR","ShortName":"ShortName","Quantity":1}]}';

      expect(actual, match);
    });
  });

  group("StockItem", () {
    test("toJson single object", () {
      final obj = FakeItem();
      final actual = jsonEncode(obj);
      final match =
          '{"ItemKey":"0`0","ImageKey":"0`0","OwnerKey":"0`0","Expires":"2020-06-06T00:00:00.000","Price":10,"Tags":["tag","tag"],"Location":"Somewhere","Views":0,"OwnerHistory":{},"EstimateValue":15,"Currency":"ZAR","ShortName":"ShortName","Quantity":1}';

      expect(actual, match);
    });

    test("toJson list", () {
      final lst = new List<StockItem>();
      lst.add(FakeItem());

      final actual = jsonEncode(lst);
      final match =
          '[{"ItemKey":"0`0","ImageKey":"0`0","OwnerKey":"0`0","Expires":"2020-06-06T00:00:00.000","Price":10,"Tags":["tag","tag"],"Location":"Somewhere","Views":0,"OwnerHistory":{},"EstimateValue":15,"Currency":"ZAR","ShortName":"ShortName","Quantity":1}]';

      expect(actual, match);
    });
  });
}

StockItem FakeItem() {
  var itemKey = new Key("0`0");
  var shortname = "ShortName";
  var imageKey = new Key("0`0");
  var ownerKey = new Key("0`0");
  var expires = new DateTime(2020, 06, 06);
  var currency = "ZAR";
  num price = 10;
  num estimate = 15;
  var tags = ["tag", "tag"];
  var location = "Somewhere";
  var history = new Map<DateTime, Key>();
  return new StockItem(itemKey, shortname, imageKey, ownerKey, expires,
      currency, price, estimate, tags, location, 0, history, 1);
}
