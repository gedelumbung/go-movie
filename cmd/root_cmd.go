package cmd

import (
	"github.com/gedelumbung/go-movie/api"
	"github.com/gedelumbung/go-movie/component"
	"github.com/gedelumbung/go-movie/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFile = ""

var rootCmd = cobra.Command{
	Use:  "movie",
	Long: "Simple Movie API with GoLang",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, serve)
	},
}

func RootCmd() *cobra.Command {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "The configuration file")
	rootCmd.AddCommand(&serveCmd)
	return &rootCmd
}

func execWithConfig(cmd *cobra.Command, fn func(config *conf.Configuration)) {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}

	fn(config)
}

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start web server",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, serve)
	},
}

func serve(config *conf.Configuration) {
	db, err := component.GetDatabaseConnection(config)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}

	log := component.GetLogger(config)

	api := api.NewAPI(config, db, log)
	api.ListenAndServe()
}
