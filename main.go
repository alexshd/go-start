package main

import (
	"github.com/shdlabs/go-start/project"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := project.NewRootCmd().Execute(); err != nil {
		logrus.Fatal(err)
	}
}
