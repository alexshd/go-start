package create_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/shdlabs/go-start/create"
	"github.com/stretchr/testify/suite"
)

type fixtureTMPFile struct {
	suite.Suite
	data    *create.TempData
	name    string
	tp      string
	wantOut string
}

func TestSute(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(fixtureTMPFile))
}

func (f *fixtureTMPFile) SetupTest() {
	f.name = "test1"
	f.tp = "{{ .Name }} - {{ .Name | ToLower }} - {{ .Name | Title -}}"
	f.data = &create.TempData{
		Name: "Cool thing", FuncMap: template.FuncMap{
			"ToLower": strings.ToLower,
			"Title":   strings.Title,
		},
	}

	f.wantOut = "Cool thing - cool thing - Cool Thing"
}

func (f *fixtureTMPFile) TestTempPopulate() {
	out := &bytes.Buffer{}
	err := create.TempPopulate(out, f.tp, f.data)
	f.NoError(err)
	f.Equal(f.wantOut, out.String())
}

func ExampleToptalURIBulder() {
	fmt.Println(create.ToptalURIBulder("go", "vscode", "macos"))
	// Output: https://www.toptal.com/developers/gitignore/api/go,vscode,macos
}
