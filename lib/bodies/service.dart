import 'package:mango_stock/bodies/stockitem.dart';

class Service extends StockItem {
  final String Url;

  Service(this.Url)
      : super(super.ImageKey, super.OwnerKey, super.Expires, super.Price,
            super.Tags, super.Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["Url"] = this.Url;

    return result;
  }
}
