package cmd

import (
	"fmt"
	"os"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	// Config holds the global configuration for kkpctl
	Config config.Config

	configPath string
	outputType string
	sortBy     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kkpctl [(-o|--output=)text|json|yaml] [(--sort=name|date)]",
	Short: "A CLI for interacting with Kubermatic Kubernetes Platform.",
	Long:  `This is a CLI for interacting with the REST API of Kubermatic Kubernetes Platform.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var err error
	config.ConfigPath = configPath
	Config, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Failed to find home directory: " + err.Error())
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&configPath, "config", home+"/.config/kkpctl/config.yaml", "The Path to the configuration file")

	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "text", "The output type to use")
	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"text", "json", "yaml"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().StringVar(&sortBy, "sort", "name", "Sort text output by which attribute (\"name\" or \"date\")")
	rootCmd.RegisterFlagCompletionFunc("sort", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "date"}, cobra.ShellCompDirectiveDefault
	})
}
