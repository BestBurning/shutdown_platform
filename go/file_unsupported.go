// +build !darwin,!linux,!windows

package shutdown_platform

import (
	"github.com/pkg/errors"
)

func shutdownPlatform(second string) (res string, err error) {
	return "", errors.New("platform unsupported")
}
