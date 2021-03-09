package main_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fixtureStart struct {
	suite.Suite
}

func TestSute(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(fixtureStart))
}

func (f *fixtureStart) SetupTest() {
}

func (f *fixtureStart) TestStart() {
	f.Equal(1, 1, "test to fail")
}
