package main

import (
	"log"

	"github.com/shdlabs/go-start/project"
)

func main() {
	if err := project.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
