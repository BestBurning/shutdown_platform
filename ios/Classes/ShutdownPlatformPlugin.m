#import "ShutdownPlatformPlugin.h"
#if __has_include(<shutdown_platform/shutdown_platform-Swift.h>)
#import <shutdown_platform/shutdown_platform-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "shutdown_platform-Swift.h"
#endif

@implementation ShutdownPlatformPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftShutdownPlatformPlugin registerWithRegistrar:registrar];
}
@end
