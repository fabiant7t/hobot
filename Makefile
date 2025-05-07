build:
	go build -o hobotdev

release: clean
	goreleaser release --clean

clean:
	rm -rf dist

docs:
	rm -rf docs
	go run main.go docs

.PHONY: docs
