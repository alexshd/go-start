// Package create data and functions handles creation
package create

import (
	_ "embed" // the way embed works with strings
	"io"
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

func TempPopulate(w io.Writer, tp string, name string) error {
	data := &tempData{
		Name: name,
		funcMap: template.FuncMap{
			"ToLower": strings.ToLower,
			"Title":   strings.Title,
		},
	}

	tmpl := template.Must(template.New("testfile").Funcs(data.funcMap).Parse(tp))
	err := tmpl.Execute(w, data)

	return errors.Wrap(err, "template execution")
}
