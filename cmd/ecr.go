package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/applicreation/aws-cli/settings"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/spf13/cobra"
)

func ecrCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ecr",
		Short: "Amazon ECR Registry",
	}

	// batch-check-layer-availability
	// batch-delete-image
	// batch-get-image
	// complete-layer-upload
	// create-repository
	// delete-lifecycle-policy
	// delete-repository
	// delete-repository-policy
	// describe-image-scan-findings
	// describe-images
	// describe-repositories
	// get-authorization-token
	// get-download-url-for-layer
	// get-lifecycle-policy
	// get-lifecycle-policy-preview
	// get-login
	cmd.AddCommand(ecrGetLoginPasswordCmd(options))
	// get-repository-policy
	// initiate-layer-upload
	// list-images
	// list-tags-for-resource
	// put-image
	// put-image-scanning-configuration
	// put-image-tag-mutability
	// put-lifecycle-policy
	// set-repository-policy
	// start-image-scan
	// start-lifecycle-policy-preview
	// tag-resource
	// untag-resource
	// upload-layer-part
	// wait

	return cmd
}

func ecrGetLoginPasswordCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-login-password",
		Run: func(cmd *cobra.Command, args []string) {
			svc := ecr.New(MakeAwsConfig(options))

			input := &ecr.GetAuthorizationTokenInput{}

			request := svc.GetAuthorizationTokenRequest(input)

			result, err := request.Send(context.TODO())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			data := result.AuthorizationData[0]

			token := data.AuthorizationToken

			authToken, _ := base64.StdEncoding.DecodeString(*token)

			password := strings.Split(string(authToken), ":")
			fmt.Printf("%+v\n", password[1])
		},
	}

	return cmd
}
