package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/create"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := realMain(); err != nil {
		logrus.Fatalf("%v", err)
	}
}

// handleName
// 1. short name => error
// 2. long name 'github.com/shdlabs/newpackage => dir `newpackage` , mod init `full`
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

func realMain() error {
	packageName := ""

	flag.StringVar(&packageName, "name", "", "-name <name>; name should match |^[a-z0-9]{2,}$|")
	flag.Parse()

	fullName := packageName

	if strings.ContainsRune(packageName, '/') {
		s := strings.Split(packageName, "/")
		packageName = s[len(s)-1]
	}

	if err := handleName(packageName); err != nil {
		flag.Usage()

		return errors.Wrap(err, "regex roles")
	}

	if err := os.Mkdir(packageName, 0754); err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	_ = os.Chdir(packageName) // just created the dir

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
		`git add .`,
		"go mod tidy",
	}

	for _, cmd := range commands {
		out, err := cmdFactory(cmd)
		if err != nil {
			return errors.Wrap(err, cmd)
		}

		if len(out) > 0 {
			logrus.Printf("%s\n", out)
		}
	}

	return nil
}

func cmdFactory(exeCommand string) ([]byte, error) {
	args := strings.Split(exeCommand, " ")

	return exec.Command(args[0], args[1:]...).Output()
}
