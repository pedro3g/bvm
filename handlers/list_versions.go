package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pedro3g/bvm/utils"
)

type Version struct {
	TagName     string    `json:"tag_name"`
	PublishedAt time.Time `json:"published_at"`
}

func ListVersions(echo bool) ([]Version, error) {
	releasesUrl := "https://api.github.com/repos/oven-sh/bun/releases"

	resp, err := http.Get(releasesUrl)

	if err != nil {
		if !echo {
			return nil, err
		}

		log.Fatalln("Could not fetch the list of releases")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		if !echo {
			return nil, err
		}

		log.Fatalln("Could not read the list of releases")
	}

	var version []Version

	err = json.Unmarshal(body, &version)

	if err != nil {
		if !echo {
			return nil, err
		}

		log.Fatalln("Could not unmarshal the JSON response")
	}

	utils.Reverse(&version)

	if echo {
		for _, version := range version {
			v := strings.Split(version.TagName, "-")[1]

			fmt.Println(v)
		}
	}

	return version, nil
}
