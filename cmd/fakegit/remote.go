package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"
)

type remoteArgs struct {
	RootArgs rootArgs `cli:"remote,subcmd"`
	Verbose  bool     `cli:"-v,--verbose" usage:"be verbose"`
}

func (_ remoteArgs) Description() string {
	return "manage set of tracked repositories"
}

func (_ remoteArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Manage the set of repositories ("remotes") whose branches you track.
`)
}

func (_ remoteArgs) ExtendedUsage_Verbose() string {
	return strings.TrimSpace(`
Be a little more verbose and show remote url after name. NOTE: This must be
placed between remote and subcommand.
`)
}

func remote(ctx context.Context, args remoteArgs) error {
	return json.NewEncoder(os.Stdout).Encode(args)
}

type remoteAddArgs struct {
	RemoteArgs remoteArgs `cli:"add,subcmd"`
	Track      string     `cli:"-t" value:"branch" usage:"track only the given branch"`
	Master     string     `cli:"-m" value:"master" usage:"set HEAD to track the given branch"`
	Fetch      bool       `cli:"-f" usage:"run fetch after adding remote"`
	Tags       bool       `cli:"--tags" usage:"import every tag from remote"`
	NoTags     bool       `cli:"--no-tags" usage:"do not import every tag from remote"`
	Name       string     `cli:"name"`
	URL        string     `cli:"url"`
}

func (_ remoteAddArgs) Description() string {
	return "add a remote for a repository"
}

func (_ remoteAddArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Adds a remote named <name> for the repository at <url>. The command git fetch
<name> can then be used to create and update remote-tracking branches
<name>/<branch>.
`)
}

func (_ remoteAddArgs) ExtendedUsage_Track() string {
	return strings.TrimSpace(`
Instead of the default glob refspec for the remote to track all branches under
the refs/remotes/<name>/ namespace, a refspec to track only <branch> is created.
You can give more than one -t <branch> to track multiple branches without
grabbing all branches.
`)
}

func (_ remoteAddArgs) ExtendedUsage_Master() string {
	return strings.TrimSpace(`
A symbolic-ref refs/remotes/<name>/HEAD is set up to point at remote's <master>
branch. See also the set-head command.
`)
}

func (_ remoteAddArgs) ExtendedUsage_Fetch() string {
	return strings.TrimSpace(`
git fetch <name> is run immediately after the remote information is set up.
`)
}

func (_ remoteAddArgs) ExtendedUsage_Tags() string {
	return strings.TrimSpace(`
git fetch <name> imports every tag from the remote repository.
`)
}

func (_ remoteAddArgs) ExtendedUsage_NoTags() string {
	return strings.TrimSpace(`
git fetch <name> does not import tags from the remote repository.
`)
}

func remoteAdd(ctx context.Context, args remoteAddArgs) error {
	return json.NewEncoder(os.Stdout).Encode(args)
}

type remoteSetURLArgs struct {
	RemoteArgs remoteArgs `cli:"set-url,subcmd"`
	Push       bool       `cli:"--push" usage:"push URLs are manipulated instead of fetch URLs"`
	Add        bool       `cli:"--add" usage:"instead of changing existing URLs, new URL is added"`
	Delete     bool       `cli:"--delete" usage:"all URLs matching regex <url> are deleted for remote <name>"`
	Name       string     `cli:"name"`
	NewURL     string     `cli:"newurl"`
	OldURL     []string   `cli:"oldurl..."`
}

func (_ remoteSetURLArgs) Description() string {
	return "change URLs for the remote"
}

func (_ remoteSetURLArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Changes URLs for the remote. Sets first URL for remote <name> that matches regex
<oldurl> (first URL if no <oldurl> is given) to <newurl>. If <oldurl> doesn't
match any URL, an error occurs and nothing is changed.

Note that the push URL and the fetch URL, even though they can be set
differently, must still refer to the same place. What you pushed to the push URL
should be what you would see if you immediately fetched from the fetch URL. If
you are trying to fetch from one place (e.g. your upstream) and push to another
(e.g. your publishing repository), use two separate remotes.
`)
}

func (_ remoteSetURLArgs) ExtendedUsage_Push() string {
	return strings.TrimSpace(`
Push URLs are manipulated instead of fetch URLs.
`)
}

func (_ remoteSetURLArgs) ExtendedUsage_Add() string {
	return strings.TrimSpace(`
Instead of changing existing URLs, new URL is added.
`)
}

func (_ remoteSetURLArgs) ExtendedUsage_Delete() string {
	return strings.TrimSpace(`
Instead of changing existing URLs, all URLs matching regex <url> are deleted for
remote <name>. Trying to delete all non-push URLs is an error.
`)
}

func remoteSetURL(ctx context.Context, args remoteSetURLArgs) error {
	return json.NewEncoder(os.Stdout).Encode(args)
}
