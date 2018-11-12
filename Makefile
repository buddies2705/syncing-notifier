export GO111MODULE=on

build:
	cd cmd && go build -o syncing-notifier

release:
	goreleaser --skip-publish

.PHONY: build release
