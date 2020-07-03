class Tag {
  final String Description;

  Tag(this.Description);

  Map<String, dynamic> toJson() {
    return {"Description": this.Description};
  }
}
