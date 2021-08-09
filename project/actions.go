// Package project commands
package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/create"
)

type action func(string) error

func execute(name string, act ...action) error {
	var err error

	for _, f := range act {
		if err != nil {
			break
		}

		err = f(name)
	}

	return errors.Wrap(err, "execute error")
}

func verify(name string) error {
	r := regexp.MustCompile(`^[a-z0-9]{2,}$`)
	if !r.MatchString(name) {
		return nameError
	}

	return nil
}

func mkdir(name string) error {
	return errors.Wrap(
		os.Mkdir(name, 0754),
		"failed to create directory",
	)
}

func chdir(name string) error {
	return os.Chdir(name)
}

func mktest(name string) error {
	f, _ := os.Create(fmt.Sprintf("%s_test.go", name))

	return errors.Wrap(
		create.TempPopulate(f, create.TempTestFile, name),
		"failed to create test file",
	)
}

func mkgomod(name string) error {
	f, _ := os.Create("go.mod")

	return errors.Wrap(
		create.TempPopulate(f, create.TempModFile, name),
		"failed to create test file",
	)
}

func gitInit(string) error {
	_, err := git.PlainInit(filepath.Join(".", git.GitDirName), true)

	return err
}

func gitAddCommit(string) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	w, _ := r.Worktree()

	if err = w.AddGlob("."); err != nil {
		return err
	}

	_, err = w.Commit("initial commit", &git.CommitOptions{All: true})

	return err
}

func runbash(string) error {
	cmd := exec.Command("go", "mod", "tidy")

	return cmd.Run()
}

//go:generate stringer -type=nameErrors
type nameErrors int

const (
	nameError nameErrors = iota
)

func (e nameErrors) Error() string {
	return e.String()
}
