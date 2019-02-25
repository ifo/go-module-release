package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	version, err := VersionString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}

func VersionString() (string, error) {
	out, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return "", err
	}
	return strings.Split(string(out), "-")[0], nil
}
