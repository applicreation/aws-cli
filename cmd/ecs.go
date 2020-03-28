package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/applicreation/aws-cli/settings"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/spf13/cobra"
)

func ecsCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ecs",
		Short: "Amazon Elastic Container Service",
	}

	// create-capacity-provider
	// create-cluster
	// create-service
	// create-task-set
	// delete-account-setting
	// delete-attributes
	// delete-cluster
	// delete-service
	// delete-task-set
	// deploy
	// deregister-container-instance
	// deregister-task-definition
	// describe-capacity-providers
	// describe-clusters
	// describe-container-instances
	// describe-services
	// describe-task-definition
	// describe-task-sets
	// describe-tasks
	// discover-poll-endpoint
	// list-account-settings
	// list-attributes
	// list-clusters
	// list-container-instances
	// list-services
	// list-tags-for-resource
	// list-task-definition-families
	// list-task-definitions
	// list-tasks
	// put-account-setting
	// put-account-setting-default
	// put-attributes
	// put-cluster-capacity-providers
	// register-container-instance
	// register-task-definition
	// run-task
	// start-task
	// stop-task
	// submit-attachment-state-changes
	// submit-container-state-change
	// submit-task-state-change
	// tag-resource
	// untag-resource
	// update-cluster-settings
	// update-container-agent
	// update-container-instances-state
	cmd.AddCommand(ecsUpdateServiceCmd(options))
	// update-service-primary-task-set
	// update-task-set
	// wait

	return cmd
}

func ecsUpdateServiceCmd(options *settings.Options) *cobra.Command {
	var Cluster string
	var Service string
	var ForceNewDeployment bool

	cmd := &cobra.Command{
		Use: "update-service",
		Run: func(cmd *cobra.Command, args []string) {
			svc := ecs.New(MakeAwsConfig(options))

			input := &ecs.UpdateServiceInput{
				Cluster:            aws.String(Cluster),
				Service:            aws.String(Service),
				ForceNewDeployment: aws.Bool(ForceNewDeployment),
			}

			request := svc.UpdateServiceRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&Cluster, "cluster", "", "")
	cmd.Flags().StringVar(&Service, "service", "", "")
	cmd.Flags().BoolVar(&ForceNewDeployment, "force-new-deployment", false, "")

	cmd.MarkFlagRequired("service")

	return cmd
}
