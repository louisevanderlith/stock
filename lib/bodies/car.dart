import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_stock/bodies/tag.dart';
import 'package:mango_ui/keys.dart';

class Car extends StockItem {
  final Key VehicleKey;
  final String Info;
  final num Year;
  final num Mileage;
  final bool HasNatis;
  final BigInt EstValue;
  final DateTime LicenseExpiry;

  Car(
      this.VehicleKey,
      this.Info,
      this.Year,
      this.Mileage,
      this.HasNatis,
      this.EstValue,
      this.LicenseExpiry,
      Key ImageKey,
      Key EntityKey,
      DateTime Expires,
      int Price,
      List<Tag> Tags,
      String Location)
      : super(ImageKey, EntityKey, Expires, Price, Tags, Location);

  Map<String, dynamic> toJson() {
    final result = super.toJson();
    result["VehicleKey"] = this.VehicleKey;
    result["Info"] = this.Info;
    result["Year"] = this.Year;
    result["Mileage"] = this.Mileage;
    result["HasNatis"] = this.HasNatis;
    result["EstValue"] = this.EstValue;
    result["LicenseExpiry"] = this.LicenseExpiry;

    return result;
  }
}
