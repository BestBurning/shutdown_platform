import 'dart:async';

import 'package:flutter/services.dart';

class ShutdownPlatform {
  static const MethodChannel _channel =
      const MethodChannel('com.di1shaui.flutter/shutdown_platform');


  static Future<String> get shutdown async {
    final String str = await _channel.invokeMethod('shutdown');
    return str;
  }

}
