package main

import (
	"fmt"
	"github.com/jplu/visual-search-backend/cmd"
	"github.com/jplu/visual-search-backend/implem/ann.faiss"
	"github.com/jplu/visual-search-backend/implem/features.image"
	"github.com/jplu/visual-search-backend/implem/router.gin"
	"github.com/jplu/visual-search-backend/infra"
	"github.com/jplu/visual-search-backend/uc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:   "visual-search-backend",
	Short: "visual-search-backend",
	Run:   func(cmd *cobra.Command, args []string) { run() }, // nolint: unparam
}

// Version command that will display the build number and version (if any)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) }, // nolint: unparam
}

func main() {
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddAllFlags(rootCmd)
	rootCmd.AddCommand(versionCmd)

	// Run the command
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't start")
	}
}

func run() {
	logrus.WithField("version", Version).WithField("build", Build).Info("Starting")

	// Setup the server with cors and stuff
	s := infra.NewServer(
		viper.GetString("server.host"),
		viper.GetInt("server.port"),
		viper.GetString("server.mode"),
		infra.NewCorsConfig(
			viper.GetBool("server.cors.disabled"),
			viper.GetBool("server.cors.all"),
			viper.GetStringSlice("server.cors.origins"),
			viper.GetStringSlice("server.cors.methods"),
			viper.GetStringSlice("server.cors.headers"),
			viper.GetStringSlice("server.cors.expose"),
		),
	)

	var err error
	var imageModelClient uc.Image

	if imageModelClient, err = image.NewImageModelClient(viper.GetString("img.host"), viper.GetInt("img.port")); imageModelClient == nil {
		logrus.WithError(err).Fatal("Couldn't connect to the Image Model gRPC server")
	}

	var faissClient uc.Ann

	if faissClient, err = faiss.NewFaissClient(viper.GetString("ann.host"), viper.GetInt("ann.port")); faissClient == nil {
		logrus.WithError(err).Fatal("Couldn't connect to the ANN gRPC server")
	}

	i := uc.NewInteractor(
		imageModelClient,
		faissClient,
	)

	r := router.NewRouter(
		i,
		s.Router,
	)

	r.SetRoutes()
	s.Start()
}
