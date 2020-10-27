import 'package:mango_ui/keys.dart';

class StockItem {
  final Key ImageKey;
  final Key OwnerKey;
  final DateTime Expires;
  final int Price;
  final List<String> Tags;
  final String Location;

  StockItem(this.ImageKey, this.OwnerKey, this.Expires, this.Price, this.Tags,
      this.Location);

  Map<String, dynamic> toJson() {
    return {
      "ImageKey": this.ImageKey,
      "OwnerKey": this.OwnerKey,
      "Expires": this.Expires,
      "Price": this.Price,
      "Tags": this.Tags,
      "Location": this.Location
    };
  }
}
