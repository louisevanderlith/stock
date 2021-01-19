import 'package:mango_ui/keys.dart';

class Product {
  final Key categoryKey;
  final List<Key> itemKeys;
  final String shortname;
  final Key imageKey;
  final DateTime expires;
  final String currency;
  final double price;
  final double estimate;
  final List<String> tags;
  final String location;
  final int views;

  Product(
      this.categoryKey,
      this.itemKeys,
      this.shortname,
      this.imageKey,
      this.expires,
      this.currency,
      this.price,
      this.estimate,
      this.tags,
      this.location,
      this.views);

  Map<String, dynamic> toJson() {
    return {
      "CategoryKey": categoryKey,
      "ItemKeys": itemKeys,
      "ImageKey": imageKey,
      "Expires": expires.toIso8601String(),
      "Price": price,
      "Tags": tags,
      "Location": location,
      "Views": views,
      "EstimateValue": estimate,
      "Currency": currency,
      "ShortName": shortname,
    };
  }
}
