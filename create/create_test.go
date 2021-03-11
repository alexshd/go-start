package create_test

import (
	"bytes"
	"testing"

	"github.com/shdlabs/go-start/create"
	"github.com/stretchr/testify/suite"
)

type fixtureTMPFile struct {
	suite.Suite
	name    string
	tp      string
	wantOut string
}

func TestSute(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(fixtureTMPFile))
}

func (f *fixtureTMPFile) SetupTest() {
	f.name = "Cool thing"
	f.tp = "{{ .Name }} - {{ .Name | ToLower }} - {{ .Name | Title -}}"

	f.wantOut = "Cool thing - cool thing - Cool Thing"
}

func (f *fixtureTMPFile) TestTempPopulate() {
	out := &bytes.Buffer{}
	err := create.TempPopulate(out, f.tp, f.name)
	f.NoError(err)
	f.Equal(f.wantOut, out.String())
}
