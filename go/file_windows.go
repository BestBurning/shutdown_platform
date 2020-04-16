package shutdown_platform

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

func shutdownPlatform(second string) (res string, err error) {
	fmt.Println("将在", second, "s后关机")
	arg := []string{"-s", "-t", second }
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("关机失败")
		return "", errors.Wrap(err, "failed to shutdown")
	}
	return string(d), nil

}
