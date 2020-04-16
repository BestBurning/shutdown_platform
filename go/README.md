# shutdown_platform

This Go package implements the host-side of the Flutter [shutdown_platform](https://github.com/BestBurning/shutdown_platform) plugin.

## Usage

Import as:

```go
import shutdown_platform "github.com/BestBurning/shutdown_platform/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&shutdown_platform.ShutdownPlatformPlugin{}),
```
