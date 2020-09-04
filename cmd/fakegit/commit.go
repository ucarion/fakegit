package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"
)

type commitArgs struct {
	RootArgs          rootArgs `cli:"commit,subcmd"`
	All               bool     `cli:"-a,--all" usage:"automatically stage files"`
	Patch             bool     `cli:"-p,--patch" usage:"use interactive patch selection"`
	ReuseMessage      string   `cli:"-C,--reuse-message" value:"commit" usage:"reuse log message from commit"`
	ReeditMessage     string   `cli:"-c,--reedit-message" value:"commit" usage:"like -C, but then edit the message"`
	Fixup             string   `cli:"--fixup" value:"commit" usage:"construct fixup!-prefixed commit message for use with rebase --autosquash"`
	Squash            string   `cli:"--squash" value:"commit" usage:"construct squash!-prefixed commit message for use with rebase --autosquash"`
	ResetAuthor       bool     `cli:"--reset-author" usage:"change authorship to committer"`
	Short             bool     `cli:"--short" usage:"when doing a dry-run, give output in short-format"`
	Branch            bool     `cli:"--branch" usage:"show branch and tracking info even in short-format"`
	Porcelain         bool     `cli:"--porcelain" usage:"when doing a dry-run, output in porcelain-ready format"`
	Long              bool     `cli:"--long" usage:"when doing a dry-run, give output in long-format"`
	Null              bool     `cli:"-z,--null" usage:"print filenames verbatim, terminated with NUL"`
	File              string   `cli:"-F,--file" value:"file" usage:"take commit message from the given file"`
	Author            string   `cli:"--author" value:"author" usage:"override the commit author"`
	Date              string   `cli:"--date" value:"date" usage:"override author date used in commit"`
	Message           string   `cli:"-m,--message" value:"msg" usage:"use the given <msg> as the commit message"`
	Template          string   `cli:"-t,--template" value:"file" usage:"start editor with contents of given file when editing message"`
	Signoff           bool     `cli:"-s,--signoff" usage:"add Signed-off-by line by committer to message"`
	NoVerify          bool     `cli:"-n,--no-verify" usage:"bypass pre-commit and commit-msg hooks"`
	AllowEmpty        bool     `cli:"--allow-empty" usage:"skip check for empty commits"`
	AllowEmptyMessage bool     `cli:"--allow-empty-message" usage:"skip check for empty commit messages"`
	Cleanup           string   `cli:"--cleanup" value:"mode" usage:"determine how commit message should be cleaned up before committing"`
	Edit              bool     `cli:"-e,--edit" usage:"edit message from -F, -m, or -C"`
	NoEdit            bool     `cli:"--no-edit" usage:"do not launch an editor"`
	Amend             bool     `cli:"--amend" usage:"replace the tip of current branch by creating a new commit"`
	NoPostRequire     bool     `cli:"--no-post-rewrite" usage:"bypass post-rewrite hook"`
	Include           bool     `cli:"-i,--include" usage:"stage contents of given paths before making commit"`
	Only              bool     `cli:"-o,--only" usage:"disregard changes staged for paths not given"`
	UntrackedFiles    *string  `cli:"-u,--untracked-files" value:"mode" usage:"show untracked files"`
	Verbose           bool     `cli:"-v,--verbose" usage:"show unified diff between HEAD commit and what would be committed"`
	Quiet             bool     `cli:"-q,--quiet" usage:"suppress commit summary message"`
	DryRun            bool     `cli:"--dry-run" usage:"do not create commit; show paths to be committed"`
	Status            bool     `cli:"--status" usage:"include output of git-status(1) in message editor"`
	NoStatus          bool     `cli:"--no-status" usage:"disable --status"`
	GPGSign           *string  `cli:"-S,--gpg-sign" value:"keyid" usage:"GPG-sign commits"`
	Files             []string `cli:"...file"`
}

func (_ commitArgs) Description() string {
	return "record changes to the repository"
}

func (_ commitArgs) ExtendedDescription() string {
	return strings.TrimSpace(`
Create a new commit containing the current contents of the index and the given
log message describing the changes. The new commit is a direct child of HEAD,
usually the tip of the current branch, and the branch is updated to point to it
(unless no branch is associated with the working tree, in which case HEAD is
"detached" as described in git-checkout(1)).

The content to be committed can be specified in several ways:

1. by using git-add(1) to incrementally "add" changes to the index before using
the commit command (Note: even modified files must be "added");

2. by using git-rm(1) to remove files from the working tree and the index, again
before using the commit command;

3. by listing files as arguments to the commit command (without --interactive or
--patch switch), in which case the commit will ignore changes staged in the
index, and instead record the current content of the listed files (which
must already be known to Git);

4. by using the -a switch with the commit command to automatically "add" changes
from all known files (i.e. all files that are already listed in the index)
and to automatically "rm" files in the index that have been removed from the
working tree, and then perform the actual commit;

5. by using the --interactive or --patch switches with the commit command to
decide one by one which files or hunks should be part of the commit in
addition to contents in the index, before finalizing the operation. See the
"Interactive Mode" section of git- add(1) to learn how to operate these
modes.

The --dry-run option can be used to obtain a summary of what is included by any
of the above for the next commit by giving the same set of parameters (options
and paths).

If you make a commit and then find a mistake immediately after that, you can
recover from it with git reset.
`)
}

func (_ commitArgs) ExtendedUsage_All() string {
	return strings.TrimSpace(`
Tell the command to automatically stage files that have been modified and
deleted, but new files you have not told Git about are not affected.
`)
}

func (_ commitArgs) ExtendedUsage_Patch() string {
	return strings.TrimSpace(`
Use the interactive patch selection interface to chose which changes to commit.
See git-add(1) for details.
`)
}

func (_ commitArgs) ExtendedUsage_ReuseMessage() string {
	return strings.TrimSpace(`
Take an existing commit object, and reuse the log message and the authorship
information (including the timestamp) when creating the commit.
`)
}

func (_ commitArgs) ExtendedUsage_ReeditMessage() string {
	return strings.TrimSpace(`
Like -C, but with -c the editor is invoked, so that the user can further edit
the commit message.
`)
}

func (_ commitArgs) ExtendedUsage_Fixup() string {
	return strings.TrimSpace(`
Construct a commit message for use with rebase --autosquash. The commit message
will be the subject line from the specified commit with a prefix of "fixup! ".
See git-rebase(1) for details.
`)
}

func (_ commitArgs) ExtendedUsage_Squash() string {
	return strings.TrimSpace(`
Construct a commit message for use with rebase --autosquash. The commit message
subject line is taken from the specified commit with a prefix of "squash! ". Can
be used with additional commit message options (-m/-c/-C/-F). See git-rebase(1)
for details.
`)
}

func (_ commitArgs) ExtendedUsage_ResetAuthor() string {
	return strings.TrimSpace(`
When used with -C/-c/--amend options, or when committing after a conflicting
cherry-pick, declare that the authorship of the resulting commit now belongs to
the committer. This also renews the author timestamp.
`)
}

func (_ commitArgs) ExtendedUsage_Short() string {
	return strings.TrimSpace(`
When doing a dry-run, give the output in the short-format. See git-status(1) for
details. Implies --dry-run.
`)
}

func (_ commitArgs) ExtendedUsage_Branch() string {
	return strings.TrimSpace(`
Show the branch and tracking info even in short-format.
`)
}

func (_ commitArgs) ExtendedUsage_Porcelain() string {
	return strings.TrimSpace(`
When doing a dry-run, give the output in a porcelain-ready format. See
git-status(1) for details. Implies --dry-run.
`)
}

func (_ commitArgs) ExtendedUsage_Long() string {
	return strings.TrimSpace(`
When doing a dry-run, give the output in the long-format. Implies --dry-run.
`)
}

func (_ commitArgs) ExtendedUsage_Null() string {
	return strings.TrimSpace(`
When showing short or porcelain status output, print the filename verbatim and
terminate the entries with NUL, instead of LF. If no format is given, implies
the --porcelain output format. Without the -z option, filenames with "unusual"
characters are quoted as explained for the configuration variable core.quotePath
(see git-config(1)).
`)
}

func (_ commitArgs) ExtendedUsage_File() string {
	return strings.TrimSpace(`
Take the commit message from the given file. Use - to read the message from the
standard input.
`)
}

func (_ commitArgs) ExtendedUsage_Author() string {
	return strings.TrimSpace(`
Override the commit author. Specify an explicit author using the standard A U
Thor <author@example.com> format. Otherwise <author> is assumed to be a pattern
and is used to search for an existing commit by that author (i.e. rev-list --all
-i --author=<author>); the commit author is then copied from the first such
commit found.
`)
}

func (_ commitArgs) ExtendedUsage_Date() string {
	return strings.TrimSpace(`
Override the author date used in the commit.
`)
}

func (_ commitArgs) ExtendedUsage_Message() string {
	return strings.TrimSpace(`
Use the given <msg> as the commit message. If multiple -m options are given,
their values are concatenated as separate paragraphs.

The -m option is mutually exclusive with -c, -C, and -F.
`)
}

func (_ commitArgs) ExtendedUsage_Template() string {
	return strings.TrimSpace(`
When editing the commit message, start the editor with the contents in the given
file. The commit.template configuration variable is often used to give this
option implicitly to the command. This mechanism can be used by projects that
want to guide participants with some hints on what to write in the message in
what order. If the user exits the editor without editing the message, the commit
is aborted. This has no effect when a message is given by other means, e.g. with
the -m or -F options.
`)
}

func (_ commitArgs) ExtendedUsage_Signoff() string {
	return strings.TrimSpace(`
Add Signed-off-by line by the committer at the end of the commit log message.
The meaning of a signoff depends on the project, but it typically certifies that
committer has the rights to submit this work under the same license and agrees
to a Developer Certificate of Origin (see http://developercertificate.org/ for
more information).
`)
}

func (_ commitArgs) ExtendedUsage_NoVerify() string {
	return strings.TrimSpace(`
This option bypasses the pre-commit and commit-msg hooks. See also githooks(5).
`)
}

func (_ commitArgs) ExtendedUsage_AllowEmpty() string {
	return strings.TrimSpace(`
Usually recording a commit that has the exact same tree as its sole parent
commit is a mistake, and the command prevents you from making such a commit.
This option bypasses the safety, and is primarily for use by foreign SCM
interface scripts.
`)
}

func (_ commitArgs) ExtendedUsage_AllowEmptyMessage() string {
	return strings.TrimSpace(`
Like --allow-empty this command is primarily for use by foreign SCM interface
scripts. It allows you to create a commit with an empty commit message without
using plumbing commands like git-commit-tree(1).
`)
}

func (_ commitArgs) ExtendedUsage_Cleanup() string {
	return strings.TrimSpace(`
This option determines how the supplied commit message should be cleaned up
before committing. The <mode> can be strip, whitespace, verbatim, scissors or
default.

strip: Strip leading and trailing empty lines, trailing whitespace, commentary
and collapse consecutive empty lines.

whitespace: Same as strip except #commentary is not removed.

verbatim: Do not change the message at all.

scissors: Same as whitespace except that everything from (and including) the
line found below is truncated, if the message is to be edited. "#" can be
customized with core.commentChar.

		# ------------------------ >8 ------------------------

default: Same as strip if the message is to be edited. Otherwise whitespace.

The default can be changed by the commit.cleanup configuration variable (see
git-config(1)).
`)
}

func (_ commitArgs) ExtendedUsage_Edit() string {
	return strings.TrimSpace(`
The message taken from file with -F, command line with -m, and from commit
object with -C are usually used as the commit log message unmodified. This
option lets you further edit the message taken from these sources.
`)
}

func (_ commitArgs) ExtendedUsage_NoEdit() string {
	return strings.TrimSpace(`
Use the selected commit message without launching an editor. For example, git
commit --amend --no-edit amends a commit without changing its commit message.
`)
}

func (_ commitArgs) ExtendedUsage_Amend() string {
	return strings.TrimSpace(`
Replace the tip of the current branch by creating a new commit. The recorded
tree is prepared as usual (including the effect of the -i and -o options and
explicit pathspec), and the message from the original commit is used as the
starting point, instead of an empty message, when no other message is specified
from the command line via options such as -m, -F, -c, etc. The new commit has
the same parents and author as the current one (the --reset-author option can
countermand this).

It is a rough equivalent for:

			$ git reset --soft HEAD^

			$ ... do something else to come up with the right tree ...

			$ git commit -c ORIG_HEAD

but can be used to amend a merge commit.

You should understand the implications of rewriting history if you amend a
commit that has already been published. (See the "RECOVERING FROM UPSTREAM
REBASE" section in git-rebase(1).)
`)
}

func (_ commitArgs) ExtendedUsage_NoPostRequire() string {
	return strings.TrimSpace(`
Bypass the post-rewrite hook.
`)
}

func (_ commitArgs) ExtendedUsage_Include() string {
	return strings.TrimSpace(`
Before making a commit out of staged contents so far, stage the contents of
paths given on the command line as well. This is usually not what you want
unless you are concluding a conflicted merge.
`)
}

func (_ commitArgs) ExtendedUsage_Only() string {
	return strings.TrimSpace(`
Make a commit by taking the updated working tree contents of the paths specified
on the command line, disregarding any contents that have been staged for other
paths. This is the default mode of operation of git commit if any paths are
given on the command line, in which case this option can be omitted. If this
option is specified together with --amend, then no paths need to be specified,
which can be used to amend the last commit without committing changes that have
already been staged. If used together with --allow-empty paths are also not
required, and an empty commit will be created.
`)
}

func (_ commitArgs) ExtendedUsage_UntrackedFiles() string {
	return strings.TrimSpace(`
Show untracked files.

The mode parameter is optional (defaults to all), and is used to specify the
handling of untracked files; when -u is not used, the default is normal, i.e.
show untracked files and directories.

The possible options are:

o   no - Show no untracked files

o   normal - Shows untracked files and directories

o   all - Also shows individual files in untracked directories.

The default can be changed using the status.showUntrackedFiles configuration
variable documented in git-config(1).
`)
}

func (_ commitArgs) ExtendedUsage_Verbose() string {
	return strings.TrimSpace(`
Show unified diff between the HEAD commit and what would be committed at the
bottom of the commit message template to help the user describe the commit by
reminding what changes the commit has. Note that this diff output doesn't have
its lines prefixed with #. This diff will not be a part of the commit message.
See the commit.verbose configuration variable in git-config(1).

If specified twice, show in addition the unified diff between what would be
committed and the worktree files, i.e. the unstaged changes to tracked files.
`)
}

func (_ commitArgs) ExtendedUsage_Quiet() string {
	return strings.TrimSpace(`
Suppress commit summary message.
`)
}

func (_ commitArgs) ExtendedUsage_DryRun() string {
	return strings.TrimSpace(`
Do not create a commit, but show a list of paths that are to be committed, paths
with local changes that will be left uncommitted and paths that are untracked.
`)
}

func (_ commitArgs) ExtendedUsage_Status() string {
	return strings.TrimSpace(`
Include the output of git-status(1) in the commit message template when using an
editor to prepare the commit message. Defaults to on, but can be used to
override configuration variable commit.status.
`)
}

func (_ commitArgs) ExtendedUsage_NoStatus() string {
	return strings.TrimSpace(`
Do not include the output of git-status(1) in the commit message template when
using an editor to prepare the default commit message.
`)
}

func (_ commitArgs) ExtendedUsage_GPGSign() string {
	return strings.TrimSpace(`
GPG-sign commits. The keyid argument is optional and defaults to the committer
identity; if specified, it must be stuck to the option without a space.
`)
}

func commit(ctx context.Context, args commitArgs) error {
	return json.NewEncoder(os.Stdout).Encode(args)
}
