package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/applicreation/aws-cli/settings"
	"github.com/spf13/cobra"
)

var rootOptions *settings.Options

var rootCmd *cobra.Command

var awsConfig aws.Config

func Execute() {
	command := MakeCommands()
	if err := command.Execute(); err != nil {
		os.Exit(-1)
	}
}

func MakeCommands() *cobra.Command {
	// root options
	rootOptions = &settings.Options{}

	// root command
	rootCmd = &cobra.Command{
		Use: "aws-cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("usage: aws-cli [options] <command> <subcommand> [<subcommand> ...] [parameters]")
			fmt.Println()
			fmt.Println("To see help text, you can run:")
			fmt.Println()
			fmt.Println("\taws-cli help")
			fmt.Println("\taws-cli <command> help")
			fmt.Println("\taws-cli <command> <subcommand> help")
		},
	}

	rootCmd.AddCommand(acmCmd())
	rootCmd.AddCommand(stsCmd())
	rootCmd.AddCommand(versionCmd())

	// root flags
	rootCmd.PersistentFlags().StringVar(&rootOptions.Profile, "profile", "default", "")
	rootCmd.PersistentFlags().StringVar(&rootOptions.Region, "region", "", "")

	return rootCmd
}

func init() {
	cobra.OnInitialize(func() {
		cfg, err := external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(rootOptions.Profile))
		if err != nil {
			panic("failed to load config, " + err.Error())
		}

		if rootOptions.Region == "" {
			rootOptions.Region = cfg.Region
		}

		awsConfig = aws.Config{
			Region:                         rootOptions.Region,
			Credentials:                    cfg.Credentials,
			EndpointResolver:               cfg.EndpointResolver,
			HTTPClient:                     cfg.HTTPClient,
			Handlers:                       cfg.Handlers,
			Retryer:                        cfg.Retryer,
			LogLevel:                       cfg.LogLevel,
			Logger:                         cfg.Logger,
			DisableRestProtocolURICleaning: cfg.DisableRestProtocolURICleaning,
			DisableEndpointHostPrefix:      cfg.DisableEndpointHostPrefix,
			EnableEndpointDiscovery:        cfg.EnableEndpointDiscovery,
			ConfigSources:                  cfg.ConfigSources,
		}
	})
}
