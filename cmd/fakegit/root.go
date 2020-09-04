package main

type rootArgs struct {
	Paginate bool   `cli:"-p,--paginate" usage:"output into a pager"`
	Bare     bool   `cli:"--bare" usage:"treat repo as a bare repo"`
	GitDir   string `cli:"--git-dir" usage:"set the path to the repo"`
}
