import 'package:mango_stock/bodies/tag.dart';
import 'package:mango_ui/keys.dart';

class StockItem {
  final Key ImageKey;
  final Key EntityKey;
  final DateTime Expires;
  final int Price;
  final List<Tag> Tags;
  final String Location;

  StockItem(this.ImageKey, this.EntityKey, this.Expires, this.Price, this.Tags, this.Location);

  Map<String, dynamic> toJson() {
    return {
      "ImageKey": this.ImageKey,
      "EntityKey": this.EntityKey,
      "Expires": this.Expires,
      "Price": this.Price,
      "Tags": this.Tags,
      "Location": this.Location
    };
  }
}
