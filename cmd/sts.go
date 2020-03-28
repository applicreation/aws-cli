package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/applicreation/aws-cli/settings"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

func stsCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "AWS Security Token Service",
	}

	cmd.AddCommand(stsAssumeRoleCmd(options))

	return cmd
}

func stsAssumeRoleCmd(options *settings.Options) *cobra.Command {
	var RoleArn string
	var RoleSessionName string

	cmd := &cobra.Command{
		Use: "assume-role",
		Run: func(cmd *cobra.Command, args []string) {
			svc := sts.New(MakeAwsConfig(options))

			input := &sts.AssumeRoleInput{
				RoleArn:         aws.String(RoleArn),
				RoleSessionName: aws.String(RoleSessionName),
			}

			request := svc.AssumeRoleRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&RoleArn, "role-arn", "", "")
	cmd.Flags().StringVar(&RoleSessionName, "role-session-name", "", "")

	cmd.MarkFlagRequired("role-arn")
	cmd.MarkFlagRequired("role-session-name")

	return cmd
}
