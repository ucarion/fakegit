package main

import (
	"context"

	"github.com/ucarion/cli"
)

func main() {
	cli.Run(
		context.Background(),
		root,
		status,
		remote,
		remoteAdd,
		remoteSetURL,
		commit,
	)
}
