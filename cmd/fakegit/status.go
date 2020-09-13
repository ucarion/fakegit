package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"
)

type statusArgs struct {
	RootArgs  rootArgs `cli:"status,subcmd"`
	Verbose   bool     `cli:"-v,--verbose" usage:"be verbose"`
	Short     bool     `cli:"-s,--short" usage:"show status concisely"`
	Pathspecs []string `cli:"pathspec..."`
}

func (_ statusArgs) Description() string {
	return "show the working tree status"
}

func (_ statusArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Displays paths that have differences between the index file and the current HEAD
commit, paths that have differences between the working tree and the index file,
and paths in the working tree that are not tracked by Git (and are not ignored
by gitignore(5)). The first are what you would commit by running git commit; the
second and third are what you could commit by running git add before running git
commit.
`)
}

func (_ statusArgs) ExtendedUsage_Verbose() string {
	return strings.TrimSpace(`
In addition to the names of files that have been changed, also show the textual
changes that are staged to be committed (i.e., like the output of git diff
--cached). If -v is specified twice, then also show the changes in the working
tree that have not yet been staged (i.e., like the output of git diff).
`)
}

func (_ statusArgs) ExtendedUsage_Short() string {
	return "Give the output in the short-format."
}

func status(ctx context.Context, args statusArgs) error {
	return json.NewEncoder(os.Stdout).Encode(args)
}
