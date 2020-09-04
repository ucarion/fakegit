fakegit: cmd/fakegit/*.go
	go build -o fakegit ./cmd/fakegit/...

.PHONY: fakegit-man
fakegit-man: fakegit
	UCARION_CLI_GENERATE_MAN="." ./fakegit
