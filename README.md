# fakegit

`fakegit` is a dummy subset of `git`, meant to serve as demo of how you can use
[`github.com/ucarion/cli`](https://github.com/ucarion/cli) to build rich CLI
tools with options, subcommands, man pages, and autocompletion. Also included is
[a working Homebrew recipe for
`fakegit`](https://github.com/ucarion/homebrew-fakegit), which you can use as a
starting point to make Homebrew recipes for your own tools.

## Installing fakegit

Shipping CLI tools to end users can be tricky; if you're not intimiately
familiar with how shells and Unix and really work, it can feel like a bit of a
lost art. This repo includes [a set of GitHub
actions](./.github/workflows/build.yml) that compiles `fakegit` on every push,
and generates a new set of macOS, Linux, and Windows binaries on every tag. You
can use these GitHub actions as a starting point in your own projects.

To install `fakegit` on macOS, you can run:

```bash
brew install ucarion/fakegit/fakegit
```

(The Homebrew "tap" for `fakegit` lives in
[`github.com/ucarion/homebrew-fakegit`](https://github.com/ucarion/homebrew-fakegit)).

To install `fakegit` on Linux or Windows, you can download [the latest release
build](https://github.com/ucarion/fakegit/releases/latest), and then place the
`fakegit` binary within to your `PATH`. On Linux, you can additionally install
the man pages (they have the file extention `.1`) to your `MANPATH`, and to set
up Bash/Zsh autocompletion you can run (or add to your `bashrc`/`zshrc`):

```bash
complete -o bashdefault -o default -C fakegit fakegit
```

To verify the installation, you can run:

```bash
fakegit --version
```

And, on macOS or Linux:

```bash
man fakegit
```

## Using fakegit

This section is about how to use the `fakegit` tool in this repo by compiling
and running the tool from source. If you plan on using `github.com/ucarion/cli`
in your project, you'll probably have a similar developer experience when
hacking on your project.

### Running fakegit

To build `fakegit`, you just need `make` and a working Go toolchain. Then, run
`make build/fakegit` to compile `fakegit`, which you can then run as
`./build/fakegit`.

`fakegit` has essentially no functionality; all it does is print the options you
gave it. For instance, if you run:

```bash
make build/fakegit
./build/fakegit --bare status --verbose x y z | jq
```

You'll get as output:

```json
{
  "RootArgs": {
    "Version": false,
    "Paginate": false,
    "Bare": true,
    "GitDir": ""
  },
  "Verbose": true,
  "Short": false,
  "Pathspecs": [
    "x",
    "y",
    "z"
  ]
}
```

However, one piece of functionality *is* implemented. If you pass `--version` to
`fakegit`, it will print the version of the tool you installed:

```bash
./build/fakegit --version
```

```text
v0.1.14
```

This is included to show how you can have version information in your CLI tool;
`github.com/ucarion/cli` does not try to understand versions at all; instead,
`fakegit` injects a version from version control using `git describe --tags` and
`go build -ldflags`; see the `Makefile` recipe for `build/fakegit`, as well as
`cmd/fakegit/root.go`, to see how that happens.

### Building and viewing the man pages

To build the `fakegit` man pages, run:

```bash
make fakegit_man
```

Man pages for CLI tools have a file extention of `.1`. You can view the
generated man pages by running (for example):

```bash
man ./build/fakegit.1
```

Each subcommand gets its own man page.

### Setting up local Bash/Zsh autocompletion

To set up Bash or Zsh autocompletion for your locally-built `fakegit`, run:

```bash
source fakegit_completion.sh
```

Then, you should be able to run:

```bash
# Don't actually type out <TAB>; press on the TAB key on your keyboard.
./build/fakegit <TAB>
```

```text
--bare      --git-dir   --paginate  --version   commit      remote      status
```
