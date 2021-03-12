package main

import (
	"time"

	"github.com/shdlabs/go-start/config"
	"github.com/shdlabs/go-start/project"

	"github.com/sirupsen/logrus"
)

func main() {
	defer config.Measure(time.Now(), "main")

	if err := project.NewRootCmd().Execute(); err != nil {
		logrus.Fatal(err)
	}
}
