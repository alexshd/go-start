package start

import (
	"github.com/kr/pretty"
	"github.com/shdlabs/go-start/project"

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
	a.Cmd = &cobra.Command{
		Use:   "go-start",
		Short: "start or manage projects or packages",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			pretty.Println(cmd.Flag("config"))
			pretty.Println(args)
		},
	}
	// a.Cfg = new(config.Settings)

	// cobra.OnInitialize(a.Cfg.InitConfig)
	// rootCmd.PersistentFlags().StringVar(&a.Cfg.File, "config", "", "config file (default is $HOME/.go-start.yaml)")
	// a.Cmd = rootCmd
	a.Cmd.AddCommand(project.BuildProject())
	return a.Cmd
}
