package config

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Settings struct {
	File string
}

// InitConfig reads in config file and ENV variables if set.
func (cfg *Settings) InitConfig() {
	if cfg != nil {
		// Use config file from the flag.
		viper.SetConfigFile(cfg.File)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-start" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-start")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func Measure(tn time.Time, name string) {
	logrus.Infof("%s took: %v", name, time.Since(tn))
}
