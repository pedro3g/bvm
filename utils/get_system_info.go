package utils

import (
	"errors"
	"runtime"
)

func GetSystemInfo() (platform string, arch string, err error) {
	archEquivalents := map[string]string{
		"386":   "x32",
		"amd64": "aarch64",
	}

	platform = runtime.GOOS
	arch = archEquivalents[runtime.GOARCH]
	err = nil

	if arch == "" {
		err = errors.New("unable to detect your processor architecture")
	}

	return platform, arch, err
}
