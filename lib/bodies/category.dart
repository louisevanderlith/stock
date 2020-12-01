import 'package:mango_ui/keys.dart';

import 'stockitem.dart';

class Category {
  final String name;
  final String text;
  final String description;
  final String pageUrl;
  final num baseCategory;
  final String clientID;
  final Key imageKey;
  final List<StockItem> items;

  Category(this.name, this.text, this.description, this.pageUrl, this.baseCategory, this.clientID, this.imageKey, this.items);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
      "Text": text,
      "Description": description,
      "PageURL": pageUrl,
      "BaseCategory": baseCategory,
      "ClientID": clientID,
      "ImageKey": imageKey.toJson(),
      "Items": items,
    };
  }
}
