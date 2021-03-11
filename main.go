package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/create"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := realMain(); err != nil {
		logrus.Fatal(err)
	}
}

func realMain() error {
	packageName := "newpackage"

	flag.StringVar(&packageName, "name", "", "project name to create")
	flag.Parse()

	if err := os.Mkdir(packageName, 0754); err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	// Enter new directory
	if err := os.Chdir(packageName); err != nil {
		return errors.Wrap(err, "failed to enter the directory")
	}

	done := make(chan error)

	go func() { done <- mkGitingnore() }()

	f, _ := os.Create(fmt.Sprintf("%s_test.go", packageName))

	if err := create.TempPopulate(f, create.TempTestFile, packageName); err != nil {
		return errors.Wrap(err, "failed to create test file")
	}

	out, err := exec.Command("go", "mod", "init", packageName).Output()
	if err != nil {
		return errors.Wrap(err, "failed to execute `go mod init`")
	}

	logrus.Printf("%s\n", out)

	out, err = exec.Command("git", "init").Output()
	if err != nil {
		return errors.Wrap(err, "failed to execute `git init`")
	}

	logrus.Printf("%s\n", out)

	err = <-done
	if err != nil {
		return errors.Wrap(err, "failed to create gitignore file")
	}

	out, err = exec.Command("git", "add", ".").Output()
	if err != nil {
		return errors.Wrap(err, "failed to execute `git add .`")
	}

	logrus.Printf("%s\n", out)

	out, err = exec.Command("go", "mod", "tidy").Output()
	if err != nil {
		return errors.Wrap(err, "failed to execute `go mod tidy`")
	}

	logrus.Printf("%s\n", out)

	return nil
}

func mkGitingnore() error {
	res, err := http.Get(create.ToptalURIBulder("go", "vscode", "macos"))
	if err != nil {
		return errors.Wrap(err, "request  failed")
	}

	defer res.Body.Close()

	f, _ := os.Create(".gitignore")
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		return errors.Wrap(err, "failed writing to file")
	}

	return nil
}
