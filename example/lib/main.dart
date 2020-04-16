import 'package:flutter/material.dart';

import 'package:shutdown_platform/shutdown_platform.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        home: Scaffold(
            appBar: AppBar(
              title: const Text('Shutdown Platform Plugin example app'),
            ),
            body: Center(
              child: IconButton(
                icon: Icon(Icons.power_settings_new),
                onPressed: () async {
                  String str = await ShutdownPlatform.shutdown;
                  print(str);
                },
              ),
            )));
  }
}