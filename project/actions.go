// Package project commands
package project

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bitfield/script"
	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/create"
)

func handleName(name string) error {
	r := regexp.MustCompile(`^[a-z0-9]{2,}$`)
	if !r.MatchString(name) {
		return nameError
	}

	return nil
}

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

func runbash(name string) error {
	commands := []string{
		fmt.Sprintf("go mod init %s", name),
		"git init",
		"go mod tidy",
		`git add .`,
		`git commit -m'first init'`,
	}

	for _, cmd := range commands {
		_, err := script.Exec(cmd).Stdout()
		if err != nil {
			return errors.Wrapf(err, "%q", cmd)
		}
	}

	return nil
}

//go:generate stringer -type=nameErrors
// StartErrors error type.
type nameErrors int

const (
	nameError nameErrors = 11
)

func (e nameErrors) Error() string {
	return e.String()
}
