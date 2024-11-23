package main

import (
	"fmt"
	"net/http"

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
	var input string
	fmt.Println("Do you want to update? (yes/no)")
	fmt.Scanln(&input)

	if input == "yes" {
		err := doUpdate(
			"https://github.com/abatilo/distribution/releases/download/v0.0.5/distribution_darwin_arm64",
		)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Println("Update successful. Exiting.")
		return
	}
	fmt.Printf("Hello from version: %s\n", ReleaseVersion)
}
