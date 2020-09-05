package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var version string

type rootArgs struct {
	Version  bool   `cli:"--version" usage:"print version and exit"`
	Paginate bool   `cli:"-p,--paginate" usage:"output into a pager"`
	Bare     bool   `cli:"--bare" usage:"treat repo as a bare repo"`
	GitDir   string `cli:"--git-dir" usage:"set the path to the repo"`
}

func (_ rootArgs) Description() string {
	return "the stupid content tracker"
}

func (_ rootArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Git is a fast, scalable, distributed revision control system with an unusually
rich command set that provides both high-level operations and full access to
internals.

See gittutorial(7) to get started, then see giteveryday(7) for a useful minimum
set of commands. The Git User's Manual[1] has a more in-depth introduction.

After you mastered the basic concepts, you can come back to this page to learn
what commands Git offers. You can learn more about individual Git commands with
"git help command". gitcli(7) manual page gives you an overview of the
command-line command syntax.

A formatted and hyperlinked copy of the latest Git documentation can be viewed
at https://git.github.io/htmldocs/git.html or https://git-scm.com/docs.
`)
}

func (_ rootArgs) ExtendedUsage_Version() string {
	return strings.TrimSpace(`
Prints the Git suite version that the git program came from.
`)
}

func (_ rootArgs) ExtendedUsage_Pager() string {
	return strings.TrimSpace(`
Pipe all output into less (or if set, $PAGER) if standard output is a terminal.
This overrides the pager.<cmd> configuration options (see the "Configuration
Mechanism" section below).
`)
}

func root(ctx context.Context, args rootArgs) error {
	// This is the *only* actual implementation logic we'll put in fakegit:
	// supporting --version, to demonstrate how to set up injecting a version
	// into an executable.
	if args.Version {
		fmt.Println(version)
		return nil
	}

	return json.NewEncoder(os.Stdout).Encode(args)
}
