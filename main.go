package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	gitlab "github.com/xanzy/go-gitlab"
)

func isValidVersionType(vt string) bool {
	switch vt {
	case
		"major",
		"minor",
		"patch":
		return true
	}
	return false
}

func bumpVersion(projectID string, glClient *gitlab.Client, vt string) (version string, err error) {
	listOptions := &gitlab.ListTagsOptions{
		Page:    0,
		PerPage: 10,
	}

	tags, _, err := glClient.Tags.ListTags(projectID, listOptions)

	if err != nil {
		return "", err
	}

	if len(tags) == 0 {
		return "", errors.New("There are no tags yet. Create the first tag manually")
	}

	ver, err := semver.Make(tags[0].Name)

	if err != nil {
		return "", err
	}

	// updated version
	var uv string

	switch vt {
	case "major":
		uv = fmt.Sprintf("%d.0.0", (ver.Major + 1))
	case "minor":
		uv = fmt.Sprintf("%d.%d.0", ver.Major, (ver.Minor + 1))
	case "patch":
		uv = fmt.Sprintf("%d.%d.%d", ver.Major, ver.Minor, (ver.Patch + 1))
	}

	return uv, nil
}

func main() {
	if os.Args[1] != "update" {
		log.Fatal("Only update command is supported")
	}

	if len(os.Args) != 3 || isValidVersionType(os.Args[2]) != true {
		log.Fatal("The update command needs to be followed by one of the following options: major, minor, patch")
	}

	// version type
	vt := os.Args[2]

	projectID := os.Getenv("CI_PROJECT_ID")
	if projectID == "" {
		log.Fatal("CI_PROJECT_ID environment variable is not set")
	}

	gitlabToken := os.Getenv("GITLAB_TOKEN")
	if gitlabToken == "" {
		log.Fatal("GITLAB_TOKEN environment variable is not set")
	}

	glClient := gitlab.NewClient(nil, gitlabToken)

	ver, err := bumpVersion(projectID, glClient, vt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", ver)
}
