package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	_, err := git.PlainInit("./.git", true)
	if err != nil {
		logrus.Error(err)
	}
}
