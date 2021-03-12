package config

import (
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
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

	v.SetConfigName(".go-start")

	configDir, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}

	v.AddConfigPath(configDir)

	if err = v.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return errors.Wrap(err, "failed binding flags")
	}

	return errors.Wrap(v.Unmarshal(s), "failed unmarshaler")
}

func Measure(tn time.Time, name string) {
	logrus.Debugf("%s took: %v", name, time.Since(tn))
}
