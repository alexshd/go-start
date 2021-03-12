// Package project commands
package project

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/config"
	"github.com/shdlabs/go-start/create"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func BuildProject() *cobra.Command {
	defer config.Measure(time.Now(), "BuildProject")
	projectCommand := &cobra.Command{
		Use:   "project",
		Short: "manage projects",
		Long:  longHelp,

		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(makePackage(cmd.Flag("name").Value.String()))
		},
	}

	projectCommand.Flags().StringP("name", "n", "", "project name")
	// projectCommand.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	projectCommand.AddCommand(buildShow())

	return projectCommand
}

func buildShow() *cobra.Command {
	defer config.Measure(time.Now(), "buildShow")
	showCommand := &cobra.Command{
		Use:   "show",
		Short: "show info",

		Run: func(cmd *cobra.Command, args []string) {
			logrus.Println("show was called")
		},
	}

	// showCommand.Flags().StringP("name", "n", "", "project name")
	// projectCommand.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return showCommand
}

// handleName
// 3. prefix go `go-name`, `goname` => folder as given, package without prefix.
func handleName(name string) error {
	defer config.Measure(time.Now(), "handleName")

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

func makePackage(packageName string) error {
	defer config.Measure(time.Now(), "makePackage")

	fullName := packageName

	if strings.ContainsRune(packageName, '/') {
		s := strings.Split(packageName, "/")
		packageName = s[len(s)-1]
	}

	if err := handleName(packageName); err != nil {
		return errors.Wrap(err, "regex rules `^[a-z0-9]{2,}$`")
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
		"go mod tidy",
		`git add .`,
	}

	for _, cmd := range commands {
		out, err := cmdFactory(cmd)
		if err != nil {
			return errors.Wrapf(err, "%q", cmd)
		}

		if len(out) > 0 {
			logrus.Printf("%s\n", out)
		}
	}

	return nil
}

func cmdFactory(exeCommand string) ([]byte, error) {
	defer config.Measure(time.Now(), exeCommand)
	args := strings.Split(exeCommand, " ")

	return exec.Command(args[0], args[1:]...).Output()
}

const longHelp = `
Creates new Golang project:
	Name restriction: same case, at least 2 symbols, no special symbols.
	Accepts long 'github.com/example/newproj' and short 'newproj'.
	Creates:
		1. new directory ( 'newproj' in both cases )
		2. go mod init (if long name provided, with long otherwise short)
		3. creates '.gitignore' from api (go, vscode, macos)
		4. creates first test file from template.
		5. inits git repo
		6. git add .
	You should be able to run the test that should fail :)
 `
