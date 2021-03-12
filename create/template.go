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
	"time"

	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/config"
)

//go:embed test.tmpl
// TempTestFile contains the embedded string from file.
//nolint:gochecknoglobals
var TempTestFile string // embed works this way

type tempData struct {
	funcMap template.FuncMap
	Name    string
}

func TempPopulate(out io.Writer, tp string, name string) error {
	defer config.Measure(time.Now(), "TempPopulate")

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
	defer config.Measure(time.Now(), "MkGitIgnore")

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
