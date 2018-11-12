export GO111MODULE=on

build:
	cd cmd && go build -o syncing-notifier

release:
	goreleaser --rm-dist --skip-publish

.PHONY: build release
