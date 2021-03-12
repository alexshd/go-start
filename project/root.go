package project

import (
	"time"

	"github.com/shdlabs/go-start/config"
	"github.com/spf13/cobra"
)

type App struct {
	Cmd *cobra.Command
}

func NewRootCmd() *cobra.Command {
	a := new(App)

	return a.Execute()
}

func (a *App) Execute() *cobra.Command {
	defer config.Measure(time.Now(), "Execute")

	a.Cmd = &cobra.Command{
		Use:   "go-start",
		Short: "start or manage projects or packages",
	}

	a.Cmd.AddCommand(BuildProject())

	return a.Cmd
}
