import 'package:flutter/material.dart';
import 'package:notifier/Components/upsert_config_form.dart';

class Config extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Padding(
          padding: EdgeInsets.only(left: 20, right: 20, top: 15),
          child: UpsertConfigForm(),
        ),
      ],
    );
  }
}
