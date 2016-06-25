package main

import (
	"github.com/docopt/docopt-go"
	"os"
	"os/exec"
)

const (
	SUBTREE_DIR = "<subtree_dir>"
	REMOTE      = "<remote>"
	BRANCH      = "<branch>"
)

func main() {
	usage :=
		`usage:
  git-gh-deploy <subtree_dir> <remote> <branch>
  git-gh-deploy -h | --help

Options:
  -h --help     Show this screen.`

	arguments, _ := docopt.Parse(usage, nil, true, "git-gh-deploy 0.0.2", true)
	cmd := exec.Command("git", "subtree", "push", "--prefix", arguments[SUBTREE_DIR].(string), arguments[REMOTE].(string), arguments[BRANCH].(string))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
