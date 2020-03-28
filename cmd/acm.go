package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/applicreation/aws-cli/settings"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/spf13/cobra"
)

func acmCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acm",
		Short: "AWS Certificate Manager",
	}

	// add-tags-to-certificate
	cmd.AddCommand(acmDeleteCertificateCmd(options))
	cmd.AddCommand(acmDescribeCertificateCmd(options))
	// export-certificate
	cmd.AddCommand(acmGetCertificateCmd(options))
	// import-certificate
	cmd.AddCommand(acmListCertificateCmd(options))
	cmd.AddCommand(acmListTagsForCertificateCmd(options))
	// remove-tags-from-certificate
	cmd.AddCommand(acmRenewCertificateCmd(options))
	cmd.AddCommand(acmRequestCertificateCmd(options))
	cmd.AddCommand(acmResendValidationEmailCertificateCmd(options))
	// update-certificate-options
	// wait

	return cmd
}

func acmDeleteCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "delete-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.DeleteCertificateInput{
				CertificateArn: aws.String(CertificateArn),
			}

			request := svc.DeleteCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")

	cmd.MarkFlagRequired("certificate-arn")

	return cmd
}

func acmDescribeCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "describe-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.DescribeCertificateInput{
				CertificateArn: aws.String(CertificateArn),
			}

			request := svc.DescribeCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")

	cmd.MarkFlagRequired("certificate-arn")

	return cmd
}

func acmGetCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "get-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.GetCertificateInput{
				CertificateArn: aws.String(CertificateArn),
			}

			request := svc.GetCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")

	cmd.MarkFlagRequired("certificate-arn")

	return cmd
}

func acmListCertificateCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.ListCertificatesInput{}

			request := svc.ListCertificatesRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	return cmd
}

func acmListTagsForCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "list-tags-for-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.ListTagsForCertificateInput{
				CertificateArn: aws.String(CertificateArn),
			}

			request := svc.ListTagsForCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")

	cmd.MarkFlagRequired("certificate-arn")

	return cmd
}

func acmRenewCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "renew-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.RenewCertificateInput{
				CertificateArn: aws.String(CertificateArn),
			}

			request := svc.RenewCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")

	cmd.MarkFlagRequired("certificate-arn")

	return cmd
}

func acmRequestCertificateCmd(options *settings.Options) *cobra.Command {
	var DomainName string

	cmd := &cobra.Command{
		Use: "request-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.RequestCertificateInput{
				DomainName: aws.String(DomainName),
			}

			request := svc.RequestCertificateRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&DomainName, "domain-name", "", "")

	cmd.MarkFlagRequired("domain-name")

	return cmd
}

func acmResendValidationEmailCertificateCmd(options *settings.Options) *cobra.Command {
	var CertificateArn string
	var Domain string
	var ValidationDomain string

	cmd := &cobra.Command{
		Use: "resend-validation-email-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(MakeAwsConfig(options))

			input := &acm.ResendValidationEmailInput{
				CertificateArn:   aws.String(CertificateArn),
				Domain:           aws.String(Domain),
				ValidationDomain: aws.String(ValidationDomain),
			}

			request := svc.ResendValidationEmailRequest(input)

			result, err := request.Send(context.Background())
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", result)
		},
	}

	cmd.Flags().StringVar(&CertificateArn, "certificate-arn", "", "")
	cmd.Flags().StringVar(&Domain, "domain", "", "")
	cmd.Flags().StringVar(&ValidationDomain, "validation-domain", "", "")

	cmd.MarkFlagRequired("certificate-arn")
	cmd.MarkFlagRequired("domain")
	cmd.MarkFlagRequired("validation-name")

	return cmd
}
