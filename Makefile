.PHONY: version build release

VERSION := $(shell git describe --tags --dirty --match=v* 2> /dev/null || cat mod.version 2> /dev/null || echo v0)

version:
	./build/version.sh $1

build:
	VERSION=$(VERSION) ./build/build.sh

release: version build
	