import 'package:mango_ui/keys.dart';

class Category {
  final String name;
  final String text;
  final String description;
  final String pageUrl;
  final num baseCategory;
  final String clientID;
  final Key imageKey;
  final Key ownerKey;

  Category(this.name, this.text, this.description, this.pageUrl,
      this.baseCategory, this.clientID, this.imageKey, this.ownerKey);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
      "Text": text,
      "Description": description,
      "PageURL": pageUrl,
      "BaseCategory": baseCategory,
      "ClientID": clientID,
      "ImageKey": imageKey.toJson(),
      "OwnerKey": ownerKey.toJson(),
    };
  }
}
