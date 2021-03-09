package create

import (
	_ "embed" // the way embed works with strings
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//go:embed test.tmpl
// TempTestFile contains the embedded string from file.
var TempTestFile string

type TempData struct {
	FuncMap template.FuncMap
	Name    string
}

func TempPopulate(out io.Writer, tp string, data *TempData) {
	tmpl := template.Must(template.New("testfile").Funcs(data.FuncMap).Parse(tp))

	if err := tmpl.Execute(out, data); err != nil {
		logrus.Fatal(err)
	}
}

// GitIgnoreGenerator .gitignore generator.
// sending request to:
// https://www.toptal.com/developers/gitignore/api/{{.name1}},{{.name2}},...
func GitIgnoreGenerator(ignored ...string) (*http.Response, error) {
	langs := strings.Join(ignored, ",")
	formatURL := "https://www.toptal.com/developers/gitignore/api/%s"
	client := new(http.Client)

	res, err := client.Get(fmt.Sprintf(formatURL, langs))
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	return res, nil
}
