VERSION := v0.0.1
	
build:
	go build \
	  -ldflags "-X github.com/fabiant7t/hobot/cmd.Version=${VERSION}" \
	  -o hobot
