COMMANDS =
COMMANDS += build/fakegit.1
COMMANDS += build/fakegit-commit.1
COMMANDS += build/fakegit-remote.1
COMMANDS += build/fakegit-remote-add.1
COMMANDS += build/fakegit-remote-set-url.1
COMMANDS += build/fakegit-status.1

.PHONY: all
all: build/fakegit build/fakegit.sh $(COMMANDS) $(COMMANDS:.1=.html)

LDFLAGS = "-X main.version=$(shell git describe --tags)"
build/fakegit: cmd/fakegit/*.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./build/fakegit -ldflags $(LDFLAGS) ./cmd/fakegit/...

build/fakegit.sh:
	echo "complete -o bashdefault -o default -C fakegit fakegit" > build/fakegit.sh

build/%.1: fakegit_man
	:

.PHONY: fakegit_man
fakegit_man: ./build/fakegit
	UCARION_CLI_GENERATE_MAN="build" ./build/fakegit

build/%.html: build/%.1
	mandoc -Thtml $< > $@
