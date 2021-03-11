package main

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/shdlabs/go-start/shdio/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := realMain(); err != nil {
		logrus.Fatal(err)
	}
}

func realMain() error {
	srv := web.NewServer(":8000")
	logrus.Info("Starting Server on:", srv.Port)

	return errors.Wrap(http.ListenAndServe(srv.Port, srv), "failed to start")
}
