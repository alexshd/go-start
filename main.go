package main

import (
	"log"
	"time"

	"github.com/alexshd/go-start/config"
	"github.com/alexshd/go-start/project"
)

func main() {
	defer config.Measure(time.Now(), "main")

	if err := project.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
