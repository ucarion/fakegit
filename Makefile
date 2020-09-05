COMMANDS =
COMMANDS += fakegit.1
COMMANDS += fakegit-commit.1
COMMANDS += fakegit-remote.1
COMMANDS += fakegit-remote-add.1
COMMANDS += fakegit-remote-set-url.1
COMMANDS += fakegit-status.1

all: $(COMMANDS) $(COMMANDS:.1=.html)

fakegit: cmd/fakegit/*.go
	go build -o fakegit ./cmd/fakegit/...

%.1: fakegit
	UCARION_CLI_GENERATE_MAN="." ./fakegit

%.html: %.1
	mandoc -Thtml $< > $@

LDFLAGS = "-X main.version=$(shell git describe)"

fakegit_darwin:
	GOOS=darwin GOARCH=amd64 go build -o fakegit_darwin -ldflags $(LDFLAGS) ./cmd/fakegit/...

fakegit_linux:
	GOOS=linux GOARCH=amd64 go build -o fakegit_linux -ldflags $(LDFLAGS) ./cmd/fakegit/...

fakegit_windows:
	GOOS=windows GOARCH=amd64 go build -o fakegit_windows -ldflags $(LDFLAGS) ./cmd/fakegit/...
