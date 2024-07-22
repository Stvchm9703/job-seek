/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"job-seek/services/fetch_job_service/config"
	"os"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "JSFetchJob",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the fetch job service",
	Long:  `Start the fetch job service`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logLevel, _ := cmd.Flags().GetInt("verbose")

		if printConf, _ := cmd.Flags().GetBool("print"); printConf {
			config.PrintConfig()
			return
		}

		if dryrun, _ := cmd.Flags().GetBool("dry-run"); dryrun {
			pid := syscall.Getpid()
			go func() {
				time.Sleep(time.Second * time.Duration(5))
				syscall.Kill(pid, syscall.SIGTERM)
			}()
			ServerDryRun(logLevel)

			return // dry-run modes
		}

		if test, _ := cmd.Flags().GetBool("test"); test {
			ServerTestRun(logLevel)
			return // test mode

		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&config.CurrentConfigFilePath, "config", "C", "", "config file (default is $HOME/.job-seek.yaml)")
	cobra.OnInitialize(config.Setup)
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntP("verbose", "V", 0, "log message level")

	serveCmd.Flags().StringP("host", "i", "localhost", "host")
	viper.BindPFlag("host", serveCmd.Flags().Lookup("host"))

	serveCmd.Flags().IntP("port", "p", 60010, "port")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))

	serveCmd.Flags().Bool("dry-run", false, "dry-run mode with inputed command and config")
	serveCmd.Flags().Bool("test", false, "test run all related service with inputed command and config")
	serveCmd.Flags().BoolP("print", "P", false, "print config in toml format, for other service to use")

	rootCmd.AddCommand(serveCmd)
}
