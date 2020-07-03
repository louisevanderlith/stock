import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_stock/bodies/tag.dart';
import 'package:mango_ui/keys.dart';

class Property extends StockItem {
  final String Address;

  Property(this.Address, Key ImageKey, Key, EntityKey, DateTime Expires,
      int Price, List<Tag> Tags, String Location)
      : super(ImageKey, EntityKey, Expires, Price, Tags, Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["Address"] = this.Address;

    return result;
  }
}
