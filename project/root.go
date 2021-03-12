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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (a *App) Execute() *cobra.Command {
	defer config.Measure(time.Now(), "Execute")

	a.Cmd = &cobra.Command{
		Use:   "go-start",
		Short: "start or manage projects or packages",

		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) {
		// pretty.Println(cmd.Flag("config"))
		// pretty.Println(args)
		// },
	}
	// a.Cfg = new(config.Settings)

	// cobra.OnInitialize(a.Cfg.InitConfig)
	// rootCmd.PersistentFlags().StringVar(&a.Cfg.File, "config", "", "config file (default is $HOME/.go-start.yaml)")
	// a.Cmd = rootCmd
	a.Cmd.AddCommand(BuildProject())

	return a.Cmd
}
