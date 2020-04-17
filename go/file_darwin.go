package shutdown_platform

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

func shutdownPlatform(second string) (res string, err error) {
	arg := []string{"-e", `tell app "System Events" to shut down`}
	cmd := exec.Command("osascript", arg...)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("关机失败")
		fmt.Println(string(out))
		return "", errors.Wrap(err, "failed to shutdown")
	}
	return string(out), nil
}

func shutdownNow() (res string, err error) {
	arg := []string{"-e", `tell app "System Events" to shut down`}
	cmd := exec.Command("osascript", arg...)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("关机失败")
		fmt.Println(string(out))
		return "", errors.Wrap(err, "failed to shutdown")
	}
	return string(out), nil
}
