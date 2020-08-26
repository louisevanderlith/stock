import 'package:mango_stock/bodies/stockitem.dart';

class Part extends StockItem {
  final String Number;

  Part(this.Number)
      : super(super.ImageKey, super.OwnerKey, super.Expires, super.Price,
            super.Tags, super.Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["Number"] = this.Number;

    return result;
  }
}
