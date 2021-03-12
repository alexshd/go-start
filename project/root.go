package project

import (
	"time"

	"github.com/shdlabs/go-start/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return Execute()
}

func Execute() *cobra.Command {
	defer config.Measure(time.Now(), "Execute")

	rootCommand := &cobra.Command{
		Use:   "go-start",
		Short: "start or manage projects or packages",
		Args:  cobra.ExactArgs(1),
	}

	rootCommand.AddCommand(buildProject())

	return rootCommand
}
