package create_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestDirectory(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(fixtureDirectory))
}

type fixtureDirectory struct {
	suite.Suite
}

func (f *fixtureDirectory) SetupTest() {
}

func (f *fixtureDirectory) TestVerify() {
}
