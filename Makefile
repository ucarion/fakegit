COMMANDS =
COMMANDS += build/fakegit.1
COMMANDS += build/fakegit-commit.1
COMMANDS += build/fakegit-remote.1
COMMANDS += build/fakegit-remote-add.1
COMMANDS += build/fakegit-remote-set-url.1
COMMANDS += build/fakegit-status.1

.PHONY: all
all: $(COMMANDS) $(COMMANDS:.1=.html) build/fakegit_darwin build/fakegit_linux build/fakegit_windows

build/fakegit: cmd/fakegit/*.go
	go build -o ./build/fakegit ./cmd/fakegit/...

build/%.1: fakegit_man
	:

.PHONY: fakegit_man
fakegit_man: ./build/fakegit
	UCARION_CLI_GENERATE_MAN="build" ./build/fakegit

build/%.html: build/%.1
	mandoc -Thtml $< > $@

LDFLAGS = "-X main.version=$(shell git describe --tags)"

build/fakegit_darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/fakegit_darwin -ldflags $(LDFLAGS) ./cmd/fakegit/...

build/fakegit_linux:
	GOOS=linux GOARCH=amd64 go build -o build/fakegit_linux -ldflags $(LDFLAGS) ./cmd/fakegit/...

build/fakegit_windows:
	GOOS=windows GOARCH=amd64 go build -o build/fakegit_windows -ldflags $(LDFLAGS) ./cmd/fakegit/...
