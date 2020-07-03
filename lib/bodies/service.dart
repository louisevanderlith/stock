import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_stock/bodies/tag.dart';
import 'package:mango_ui/keys.dart';

class Service extends StockItem {
  final String Url;

  Service(this.Url, Key ImageKey, Key EntityKey, DateTime Expires, int Price,
      List<Tag> Tags, String Location)
      : super(ImageKey, EntityKey, Expires, Price, Tags, Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["Url"] = this.Url;

    return result;
  }
}
