package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Settings struct {
	Dest     string // root directory set by --dest or config <default ".">
	Commands []string
}

type App struct {
	Config *Settings
}

func New() *App {
	c := new(Settings)

	return &App{Config: c}
}

func (s *Settings) Load(cmd *cobra.Command) error {
	v := viper.New()
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	v.BindPFlags(cmd.Flags())
	err := v.Unmarshal(s)

	return err
}

func Measure(tn time.Time, name string) {
	logrus.Debugf("%s took: %v", name, time.Since(tn))
}
