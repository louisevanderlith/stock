import 'package:mango_ui/keys.dart';

class StockItem {
  final Key stockKey;
  final String shortname;
  final Key imageKey;
  final Key ownerKey;
  final DateTime expires;
  final String currency;
  final double price;
  final double estimate;
  final List<String> tags;
  final String location;
  final int views;
  final Map<DateTime, Key> history;

  StockItem(
      this.stockKey,
      this.shortname,
      this.imageKey,
      this.ownerKey,
      this.expires,
      this.currency,
      this.price,
      this.estimate,
      this.tags,
      this.location,
      this.views,
      this.history);

  Map<String, dynamic> toJson() {
    return {
      "StockKey": stockKey,
      "ImageKey": imageKey,
      "OwnerKey": ownerKey,
      "Expires": expires,
      "Price": price,
      "Tags": tags,
      "Location": location,
      "Views": views,
      "OwnerHistory": history,
      "EstimateValue": estimate,
      "Currency": currency,
      "ShortName": shortname,
    };
  }
}
