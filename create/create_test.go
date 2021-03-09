package create_test

import (
	"bytes"
	_ "embed"
	"net/http"
	"strings"
	"testing"
	"text/template"

	"github.com/shdlabs/go-start/create"
	"github.com/stretchr/testify/assert"
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
	f.tp = "{{ .Name }} - {{ .Name | ToLower }} - {{ .Name | Title -}} "
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
	create.TempPopulate(out, f.tp, f.data)
	f.Equal(f.wantOut, out.String())
}

func TestGitIgnoreGenerator(t *testing.T) {
	type args struct {
		ignored []string
	}
	tests := []struct {
		name      string
		args      args
		want      *http.Response
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitIgnoreGenerator(tt.args.ignored...)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
