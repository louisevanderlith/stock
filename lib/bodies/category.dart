import 'package:mango_ui/keys.dart';

import 'stockitem.dart';

class Category {
  final String clientID;
  final String text;
  final String description;
  final num baseCategory;
  final Key imageKey;
  final List<StockItem> items;

  Category(this.clientID, this.text, this.description, this.baseCategory,
      this.imageKey, this.items);

  Map<String, dynamic> toJson() {
    return {
      "ClientID": clientID,
      "Text": text,
      "Description": description,
      "BaseCategory": baseCategory,
      "ImageKey": imageKey,
      "Items": items,
    };
  }
}
