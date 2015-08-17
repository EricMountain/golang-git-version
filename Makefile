# Based on https://gist.github.com/rafecolton/6049826

REV_VAR := main.RevString
VERSION_VAR := main.VersionString
REPO_VERSION := $(shell git describe --always --dirty --tags)
REPO_REV := $(shell git rev-parse --sq HEAD)
GOBUILD_VERSION_ARGS := -ldflags "-X $(REV_VAR) $(REPO_REV) -X $(VERSION_VAR) $(REPO_VERSION)"

# then, change your `go install` statement from something like this:

#go get -x $(TARGETS)
#go install -x $(TARGETS)

# to something like this:

#go get -x $(GOBUILD_VERSION_ARGS) $(TARGETS)
#go install -x $(GOBUILD_VERSION_ARGS) $(TARGETS)

build:
	go build -x $(GOBUILD_VERSION_ARGS) version_trick.go

