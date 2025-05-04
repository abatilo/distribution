package main

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/minio/selfupdate"
)

var ReleaseVersion = "dirty"

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Downloading %s\n", url)
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	fmt.Printf("Done\n")
	if err != nil {
		return err
	}
	return err
}

func main() {
	currentVersion := semver.MustParse("v0.0.0")
	newVersion := semver.MustParse("v0.0.1")

	if newVersion.GreaterThan(currentVersion) {
		fmt.Printf("New version available: %s\n", newVersion)
	}
}
