import 'package:mango_ui/keys.dart';

import 'stockitem.dart';

class Category {
  final String clientID;
  final String text;
  final String description;
  final String pageurl;
  final num baseCategory;
  final Key imageKey;
  final List<StockItem> items;

  Category(this.clientID, this.text, this.description, this.pageurl,
      this.baseCategory, this.imageKey, this.items);

  Map<String, dynamic> toJson() {
    return {
      "ClientID": clientID,
      "Text": text,
      "Description": description,
      "PageURL": pageurl,
      "BaseCategory": baseCategory,
      "ImageKey": imageKey,
      "Items": items,
    };
  }
}
