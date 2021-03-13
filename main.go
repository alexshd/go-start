package main

import (
	"log"
	"time"

	"github.com/shdlabs/go-start/config"
	"github.com/shdlabs/go-start/project"
)

func main() {
	defer config.Measure(time.Now(), "main")

	if err := project.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
