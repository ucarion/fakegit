COMMANDS =
COMMANDS += build/fakegit.1
COMMANDS += build/fakegit-commit.1
COMMANDS += build/fakegit-remote.1
COMMANDS += build/fakegit-remote-add.1
COMMANDS += build/fakegit-remote-set-url.1
COMMANDS += build/fakegit-status.1

.PHONY: all
all: $(COMMANDS) $(COMMANDS:.1=.html)

LDFLAGS = "-X main.version=$(shell git describe --tags)"
build/fakegit: cmd/fakegit/*.go
	go build -o ./build/fakegit -ldflags $(LDFLAGS) ./cmd/fakegit/...

build/%.1: fakegit_man
	:

.PHONY: fakegit_man
fakegit_man: ./build/fakegit
	UCARION_CLI_GENERATE_MAN="build" ./build/fakegit

build/%.html: build/%.1
	mandoc -Thtml $< > $@
