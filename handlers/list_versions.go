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
	Assets      []Assets
}

type Assets struct {
	Name string
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

	var versions []Version

	err = json.Unmarshal(body, &versions)

	if err != nil {
		if !echo {
			return nil, err
		}

		log.Fatalln("Could not unmarshal the JSON response")
	}

	utils.Reverse(&versions)

	if echo {
		for _, version := range versions {
			v := strings.Split(version.TagName, "-v")[1]

			fmt.Println(v)
		}
	}

	return versions, nil
}
