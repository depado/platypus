package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Depado/platypus/cmd"
	"github.com/Depado/platypus/infra"
	"github.com/Depado/platypus/mocker"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:   "platypus",
	Short: "Platypus is a very simple mock server",
	Run:   func(cmd *cobra.Command, args []string) { run() }, // nolint: unparam
}

func run() {
	var err error

	fmt.Println(cmd.Logo)
	gin.SetMode(viper.GetString("server.mode"))
	r := gin.Default()
	corsc := infra.NewCorsConfig(
		viper.GetBool("server.cors.enable"),
		viper.GetBool("server.cors.all"),
		viper.GetStringSlice("server.cors.origins"),
		viper.GetStringSlice("server.cors.methods"),
		viper.GetStringSlice("server.cors.headers"),
		viper.GetStringSlice("server.cors.expose"),
	)
	if corsc != nil {
		logrus.Info("CORS Enabled")
		r.Use(cors.New(*corsc))
	}

	logrus.Info("Generating Routes")
	if err = mocker.GenerateRoutes(viper.GetString("mock"), r); err != nil {
		logrus.WithError(err).Fatal("Couldn't generate routes")
	}
	fmt.Println()
	logrus.Infof("Running Mock Server on %s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
	if err = r.Run(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))); err != nil {
		logrus.WithError(err).Fatal("Couldn't start router")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddAllFlags(rootCmd)
	rootCmd.AddCommand(cmd.VersionCmd)

	// Run the command
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't start")
	}
}
