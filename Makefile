fakegit: cmd/fakegit/*.go
	go build -o fakegit ./cmd/fakegit/...

%.1: fakegit
	UCARION_CLI_GENERATE_MAN="." ./fakegit

%.html: %.1
	mandoc -Thtml $< > $@
