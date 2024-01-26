package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pedro3g/bvm/utils"
)

var (
	baseDir, baseDirError = utils.GetBaseDir()
)

func InstallVersion(version *string) {
	if baseDirError != nil {
		log.Fatalln("Could not get the base directory")
	}

	if checkIfInstalled(version) {
		log.Fatalln("Version already installed. Please use the 'use' command to switch to it")
	}

	available, assetName := checkAvailability(version)

	if !available {
		log.Fatalln("Version not available")
	}

	archiveUrl := fmt.Sprintf("https://github.com/oven-sh/bun/releases/download/bun-v%s/%s", *version, assetName)

	fmt.Println(archiveUrl)
}

func checkIfInstalled(version *string) bool {
	versionPath := filepath.Join(baseDir, *version)

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func checkAvailability(version *string) (bool, string) {
	versions, err := ListVersions(false)

	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range versions {
		tagName := strings.Replace(v.TagName, "bun-v", "", 1)

		if tagName == *version {

			for _, asset := range v.Assets {
				assetPath := strings.Split(asset.Name, ".")[0]
				platform, arch, err := utils.GetSystemInfo()

				if err != nil {
					log.Fatalln(err)
				}

				findPath := fmt.Sprintf("bun-%s-%s", platform, arch)

				if assetPath == findPath {
					return true, asset.Name
				}
			}
		}
	}

	return false, ""
}
