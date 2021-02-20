package cmd

import (
	"bytes"
	"os"
	"strings"

	"github.com/enjoypi/gobsdiff/wrapper"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logLevel string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bsdiff",
	Short: "bsdiff",

	PreRunE: preRunE,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Help()
		}

		lhs, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer lhs.Close()

		rhs, err := os.Open(args[1])
		if err != nil {
			return err
		}
		defer rhs.Close()

		patch := new(bytes.Buffer)
		if err := wrapper.Diff(lhs, rhs, patch); err != nil {
			return err
		}

		cmd.Println("diff size ", patch.Len())
		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.L().Info(err.Error())
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&logLevel, "log.level", "info", "level of zap")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Bool("version", false, "show version")
	rootCmd.Flags().BoolP("dry-run", "n", false, "perform a trial run with no changes made")
}

func preRunE(cmd *cobra.Command, args []string) error {

	// use flag log.level
	var logger *zap.Logger
	var err error
	if strings.ToLower(logLevel) == "debug" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}
