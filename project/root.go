package project

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCommand := new(cobra.Command)
	rootCommand.Use = "go-start"
	rootCommand.Long = longHelp
	rootCommand.Short = "start or manage projects or packages"
	rootCommand.Args = cobra.ExactArgs(1)

	rootCommand.AddCommand(subcommand("project"))
	rootCommand.AddCommand(subcommand("package"))

	return rootCommand
}

func subcommand(sub string) *cobra.Command {
	packageCmd := new(cobra.Command)
	packageCmd.Use = sub
	packageCmd.Short = sub + " <command>"
	packageCmd.Args = cobra.ExactArgs(1)

	packageCmd.AddCommand(createSub())

	return packageCmd
}

func createSub() *cobra.Command {
	create := new(cobra.Command)
	create.Use = "create"
	create.Short = "create <name>"
	create.Aliases = []string{"make", "new"}
	create.Args = cobra.ExactArgs(1)

	create.RunE = createCmd

	return create
}

func createCmd(cmd *cobra.Command, args []string) error {
	act := []action{verify, mkdir, chdir, mktest}
	if cmd.Parent().Name() == "project" {
		act = append(act, mkgomod)
		act = append(act, runbash)
	}

	return execute(args[0], act...)
}

const longHelp = `
Creates new Golang
  Project:
    Name restriction: same case, at least 2 symbols, no special symbols.
      Creates:
	1. new directory ( 'newproj' in both cases )
	2. go mod init (if long name provided, with long otherwise short)
	3. creates first test file from template.
	4. inits git repo
	5. git add .
	6. git commit -m 'first init'

  Package:
    Same name restrictions (only short name).
      Creates:
	1. new directory
	2. test file

You should be able to run the test that should fail :)
`
