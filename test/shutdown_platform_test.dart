import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:shutdown_platform/shutdown_platform.dart';

void main() {
  const MethodChannel channel = MethodChannel('shutdown_platform');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await ShutdownPlatform.platformVersion, '42');
  });
}
