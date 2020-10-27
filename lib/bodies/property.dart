import 'package:mango_stock/bodies/stockitem.dart';

class Property extends StockItem {
  final String Address;

  Property(this.Address)
      : super(super.ImageKey, super.OwnerKey, super.Expires, super.Price,
            super.Tags, super.Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["Address"] = this.Address;

    return result;
  }
}
