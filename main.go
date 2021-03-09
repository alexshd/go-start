package main

import (
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/shdlabs/go-start/create"
	"github.com/sirupsen/logrus"
)

func main() {
	data := &create.TempData{
		Name: "continue",
		FuncMap: template.FuncMap{
			"ToLower": strings.ToLower,
			"ToTitle": strings.ToTitle,
		},
	}

	create.TempPopulate(os.Stdout, create.TempTestFile, data)

	f, err := os.Create(".gitignore")
	if err != nil {
		logrus.Fatal(err)
	}

	defer f.Close()

	res, err := create.GitIgnoreGenerator("go", "vscode", "macos")
	if err != nil {
		logrus.Fatal(err)
	}

	defer res.Body.Close()
	if _, err := io.Copy(f, res.Body); err != nil {
		logrus.Fatal(err)
	}
}
