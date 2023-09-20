import 'package:flutter/material.dart';

class WebhooksList extends StatelessWidget {
  final List<String>? webhooks;

  WebhooksList({required this.webhooks});

  @override
  Widget build(BuildContext context) {
    return webhooks?.isEmpty ?? true
        ? const Center(
            child: CircularProgressIndicator(),
          )
        : Expanded(
            child: SizedBox(
                height: 200.0,
                child: ListView.builder(
                  itemCount: webhooks!.length,
                  itemBuilder: (context, index) {
                    return ListTile(
                      title: Text(webhooks![index]),
                    );
                  },
                )),
          );
  }
}
