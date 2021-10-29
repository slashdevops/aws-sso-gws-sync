/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"time"

	"github.com/slashdevops/idp-scim-sync/internal/config"
	"github.com/slashdevops/idp-scim-sync/internal/version"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	cfg        config.Config
	reqTimeout time.Duration
)

// commands root
var rootCmd = &cobra.Command{
	Use:     "idpscimcli",
	Version: version.Version,
	Short:   "Check your  AWS Single Sing-On (SSO) / Google Workspace Groups/Users",
	Long: `
This is a Command-Line Interfaced (CLI) to help you validate and check your source and target Single Sing-On endpoints.
Check your AWS Single Sign-On (SSO) / Google Workspace Groups users and groups and validate your filters over Google Worspace users and groups.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cfg = config.New()

	cobra.OnInitialize(initConfig)

	// global configuration for commands
	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", config.DefaultDebug, "enable log debug level")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogFormat, "log-format", "f", config.DefaultLogFormat, "set the log format")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogLevel, "log-level", "l", config.DefaultLogLevel, "set the log level")
	rootCmd.PersistentFlags().DurationVarP(&reqTimeout, "timeout", "", time.Second*10, "requests timeout")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	switch cfg.LogFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.Fatal("Unsupported log format")
	}

	if cfg.Debug {
		cfg.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
	}
}
