// Package create data and functions handles creation
package create

import (
	_ "embed" // the way embed works with strings
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed test.tmpl
// TempTestFile contains the embedded string from file.
//nolint:gochecknoglobals
var TempTestFile string // embed works this way

//go:embed gomod.tmpl
// TempModFile contains the embedded string from file.
//nolint:gochecknoglobals
var TempModFile string

type tempData struct {
	funcMap template.FuncMap
	Name    string
}

func TempPopulate(out io.Writer, tp string, name string) error {
	data := &tempData{
		Name: name,
		funcMap: template.FuncMap{
			"ToLower": strings.ToLower,
			"Title":   strings.Title,
		},
	}

	tmpl := template.Must(template.New("testfile").Funcs(data.funcMap).Parse(tp))
	err := tmpl.Execute(out, data)

	return errors.Wrap(err, "template execution")
}

func MkGitingnore(toIgnore ...string) error {
	url := fmt.Sprintf(
		"https://www.toptal.com/developers/gitignore/api/%s", strings.Join(toIgnore, ","),
	)

	res, err := http.Get(url)
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
