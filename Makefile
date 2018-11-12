export GO111MODULE=on

build:
	cd cmd && go build -o syncing-notifier

.PHONY: build
