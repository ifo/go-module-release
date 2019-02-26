package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	version, err := VersionString()
	if err != nil {
		log.Fatal(err)
	}

	v, err := ParseVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func VersionString() (string, error) {
	out, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return "", err
	}
	return strings.Split(string(out), "-")[0], nil
}

var regex = regexp.MustCompile(`^v(?P<major>\d*)\.(?P<minor>\d*)\.(?P<patch>\d*)`)

func ParseVersion(s string) (Version, error) {
	captures := regex.FindStringSubmatch(s)
	if len(captures) != 4 {
		return Version{}, fmt.Errorf("non-parseable version: %s\n", s)
	}
	versions := []int{}
	for _, c := range captures[1:] {
		v, err := strconv.Atoi(c)
		if err != nil {
			return Version{}, err
		}
		versions = append(versions, v)
	}
	return Version{
		Major: versions[0],
		Minor: versions[1],
		Patch: versions[2],
	}, nil
}
