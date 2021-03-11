// Package create data and functions handles creation
package create

import (
	_ "embed" // the way embed works with strings
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed test.tmpl
// TempTestFile contains the embedded string from file.
//nolint:gochecknoglobals
var TempTestFile string // embed works this way

type TempData struct {
	FuncMap template.FuncMap
	Name    string
}

func TempPopulate(out io.Writer, tp string, data *TempData) error {
	tmpl := template.Must(template.New("testfile").Funcs(data.FuncMap).Parse(tp))
	err := tmpl.Execute(out, data)

	return errors.Wrap(err, "template execution")
}

// ToptalURIBulder adds list of ignored subjects to api uri.
//
// For full list of available subjects run:
//  $ curl https://www.toptal.com/developers/gitignore/api/list
func ToptalURIBulder(toIgnore ...string) string {
	return fmt.Sprintf(
		"https://www.toptal.com/developers/gitignore/api/%s", strings.Join(toIgnore, ","),
	)
}
