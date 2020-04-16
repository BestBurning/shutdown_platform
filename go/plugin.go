package shutdown_platform

import (
	"fmt"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.di1shaui.flutter/shutdown_platform"

const (
	methodCallshutdown = "shutdown"
)

// ShutdownPlatformPlugin implements flutter.Plugin and handles method.
type ShutdownPlatformPlugin struct{}

var _ flutter.Plugin = &ShutdownPlatformPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *ShutdownPlatformPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(methodCallshutdown, p.shutdown)
	return nil
}

func (p *ShutdownPlatformPlugin) shutdown(arguments interface{}) (reply interface{}, err error) {
	res, err := shutdownPlatform("30")
	if err != nil {
		return nil, err
	}
	fmt.Println("关机成功", res)
	return
}
