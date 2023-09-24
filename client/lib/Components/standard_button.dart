import 'package:flutter/material.dart';

class StandardButton extends StatelessWidget {
  final String text;
  final Function onPressed;

  StandardButton({required this.text, required this.onPressed});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        style: ButtonStyle(
            padding: MaterialStateProperty.all<EdgeInsets>(
                EdgeInsets.fromLTRB(15, 10, 15, 10)),
            shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(8)),
            ))),
        onPressed: () {
          onPressed();
        },
        child: Text(text));
  }
}
