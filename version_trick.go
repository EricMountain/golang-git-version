package main

import (
        "fmt"
	"os"
	"path"
)

var (
	VersionString string
	RevString     string
	versionFlag = false
	revFlag = false
)

func main() {
	progName := path.Base(os.Args[0])
	if VersionString == "" {
		VersionString = "<unknown>"
	}
	fmt.Printf("%s %s\n", progName, VersionString)
	if RevString == "" {
		RevString = "<unknown>"
	}
	fmt.Println(RevString)
}

