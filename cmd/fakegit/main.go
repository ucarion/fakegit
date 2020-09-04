package main

import (
	"context"

	"github.com/ucarion/cli"
)

func main() {
	cli.Run(
		context.Background(),
		status,
		remote,
		remoteAdd,
		remoteSetURL,
		commit,
	)
}
