package project

import (
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

// check that in project directory
// make dir
// check name
//

func buildPackage() *cobra.Command {
	packageCmd := new(cobra.Command)
	packageCmd.Use = "package"
	packageCmd.Short = "package actions"
	packageCmd.Args = cobra.ExactArgs(1)

	packageCmd.AddCommand(createPackage())

	return packageCmd
}

func createPackage() *cobra.Command {
	createPackage := new(cobra.Command)
	createPackage.Use = "create <name>"
	createPackage.Short = "create package"
	createPackage.Aliases = []string{"make", "new"}
	createPackage.Short = "package actions"
	createPackage.Args = cobra.ExactArgs(1)
	createPackage.RunE = func(cmd *cobra.Command, args []string) error {
		pretty.Println(args)
		pretty.Println(cmd.Parent().Name())

		return nil
	}

	return createPackage
}
