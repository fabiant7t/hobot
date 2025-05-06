VERSION := $(shell git tag --points-at HEAD --sort=-version:refname)
	
build:
	go build \
	  -ldflags "-X github.com/fabiant7t/hobot/cmd.Version=${VERSION}" \
	  -o hobot
