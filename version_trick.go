package mithril

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

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Print version and exit")
	flag.BoolVar(&revFlag, "rev", false, "Print git revision and exit")
}

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

