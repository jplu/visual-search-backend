package cmd

import (
	"strings"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AddLoggerFlags adds support to configure the level of the logger
func AddLoggerFlags(c *cobra.Command) {
	c.PersistentFlags().String("log.level", "debug", "one of debug, info, warn, error or fatal")
	c.PersistentFlags().String("log.format", "text", "one of text or json")
	c.PersistentFlags().Bool("log.line", true, "enable filename and line in logs")
}

// AddAnnFlags add support to configure the ANN gRPC server
func AddAnnFlags(c *cobra.Command) {
	c.PersistentFlags().String("ann.host", "localhost", "host on which the ANN gRPC server should listen")
	c.PersistentFlags().Int("ann.port", 8080, "port on which the ANN gRPC server should listen")
}

// AddImageFlags add support to configure the ANN gRPC server
func AddImageFlags(c *cobra.Command) {
	c.PersistentFlags().String("img.host", "localhost", "host on which the Image gRPC server should listen")
	c.PersistentFlags().Int("img.port", 8080, "port on which the ANN gRPC server should listen")
}

// AddServerFlags adds support to configure the server
func AddServerFlags(c *cobra.Command) {
	// Server related flags
	c.PersistentFlags().String("server.host", "localhost", "host on which the server should listen")
	c.PersistentFlags().Int("server.port", 8081, "port on which the server should listen")
	c.PersistentFlags().String("server.mode", "debug", "server mode can be either 'debug', 'test' or 'release'")

	// CORS related flags
	c.PersistentFlags().Bool("server.cors.disabled", false, "disable CORS completely")
	c.PersistentFlags().StringSlice("server.cors.methods", []string{"GET", "PUT", "POST", "DELETE", "OPTION", "PATCH"}, "array of allowed method when cors is enabled")
	c.PersistentFlags().StringSlice("server.cors.headers", []string{"Origin", "Authorization", "Content-Type"}, "array of allowed headers")
	c.PersistentFlags().StringSlice("server.cors.expose", []string{}, "array of exposed headers")
	c.PersistentFlags().StringSlice("server.cors.origins", []string{}, "array of allowed origins (overwritten if all is active)")
	c.PersistentFlags().Bool("server.cors.all", true, "defines that all origins are allowed")
}

// AddConfigurationFlag adds support to provide a configuration file on the
// command line
func AddConfigurationFlag(c *cobra.Command) {
	c.PersistentFlags().String("conf", "", "configuration file to use")
	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		logrus.WithError(err).WithField("step", "AddConfigurationFlag").Fatal("Couldn't bind flags")
	}
}

// AddAllFlags will add all the flags provided in this package to the provided
// command and will bind those flags with viper
func AddAllFlags(c *cobra.Command) {
	AddLoggerFlags(c)
	AddServerFlags(c)
	AddConfigurationFlag(c)
	AddAnnFlags(c)
	AddImageFlags(c)

	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		logrus.WithError(err).WithField("step", "AddAllFlags").Fatal("Couldn't bind flags")
	}
}


// Initialize will be run when cobra finishes its initialization
func Initialize() {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("abrico")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Configuration file
	if viper.GetString("conf") != "" {
		viper.SetConfigFile(viper.GetString("conf"))
	} else {
		viper.SetConfigName("conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/config/")
	}
	hasconf := viper.ReadInConfig() == nil

	// Set log level
	lvl := viper.GetString("log.level")
	if l, err := logrus.ParseLevel(lvl); err != nil {
		logrus.WithFields(logrus.Fields{"level": lvl, "fallback": "info"}).Warn("Invalid log level")
	} else {
		logrus.SetLevel(l)
	}

	// Set log format
	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	// Defines if logrus should display filenames and line where the log ocured
	if viper.GetBool("log.line") {
		logrus.AddHook(filename.NewHook())
	}

	// Delays the log for once the logger has been setup
	if hasconf {
		logrus.WithField("file", viper.ConfigFileUsed()).Debug("Found configuration file")
	} else {
		logrus.Warn("No configuration file found")
	}
}
