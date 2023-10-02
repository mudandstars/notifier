import 'package:flutter/material.dart';

class TextInput extends StatelessWidget {
  final TextEditingController controller;
  final String placeholder;

  TextInput({required this.controller, required this.placeholder});

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      decoration: InputDecoration(
        border: OutlineInputBorder(),
        hintText: placeholder,
        isDense: true,
        contentPadding: EdgeInsets.fromLTRB(10, 15, 10, 15),
      ),
      style: TextStyle(
        fontSize: 15.0, // Specify your font size here
      ),
    );
  }
}
