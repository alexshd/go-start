// Package project commands
package project

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/create"
)

// handleName
// 3. prefix go `go-name`, `goname` => folder as given, package without prefix.
func handleName(name string) error {
	r := regexp.MustCompile(`^[a-z0-9]{2,}$`)
	if !r.MatchString(name) {
		return NameError
	}

	return nil
}

//go:generate stringer -type=StartErrors
// StartErrors error type.
type StartErrors int

const (
	NameError StartErrors = 11
)

func (e StartErrors) Error() string {
	return e.String()
}

func makeProject(packageName string) error {
	fullName := packageName

	if strings.ContainsRune(packageName, '/') {
		s := strings.Split(packageName, "/")
		packageName = s[len(s)-1]
	}

	if err := os.Mkdir(packageName, 0754); err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	_ = os.Chdir(packageName)

	done := make(chan error)

	go func() { done <- create.MkGitingnore("go", "vscode", "macos") }()

	f, _ := os.Create(fmt.Sprintf("%s_test.go", packageName))

	if err := create.TempPopulate(f, create.TempTestFile, packageName); err != nil {
		return errors.Wrap(err, "failed to create test file")
	}

	err := <-done
	if err != nil {
		return errors.Wrap(err, "failed to create gitignore file")
	}

	commands := []string{
		fmt.Sprintf("go mod init %s", fullName),
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

func makePackage(name string) error {
	if err := os.Mkdir(name, 0754); err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	_ = os.Chdir(name)
	f, _ := os.Create(fmt.Sprintf("%s_test.go", name))

	return errors.Wrap(create.TempPopulate(f, create.TempTestFile, name), "failed to create test file")
}
