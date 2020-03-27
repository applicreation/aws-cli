package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/spf13/cobra"
)

func acmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acm",
		Short: "AWS Certificate Manager",
	}

	// add-tags-to-certificate
	cmd.AddCommand(acmDeleteCertificateCmd())
	cmd.AddCommand(acmDescribeCertificateCmd())
	// export-certificate
	cmd.AddCommand(acmGetCertificateCmd())
	// import-certificate
	cmd.AddCommand(acmListCertificateCmd())
	cmd.AddCommand(acmListTagsForCertificateCmd())
	// remove-tags-from-certificate
	cmd.AddCommand(acmRenewCertificateCmd())
	cmd.AddCommand(acmRequestCertificateCmd())
	cmd.AddCommand(acmResendValidationEmailCertificateCmd())
	// update-certificate-options
	// wait

	return cmd
}

func acmDeleteCertificateCmd() *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "delete-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmDescribeCertificateCmd() *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "describe-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmGetCertificateCmd() *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "get-certificate",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmListCertificateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmListTagsForCertificateCmd() *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "list-tags-for-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmRenewCertificateCmd() *cobra.Command {
	var CertificateArn string

	cmd := &cobra.Command{
		Use: "renew-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmRequestCertificateCmd() *cobra.Command {
	var DomainName string

	cmd := &cobra.Command{
		Use: "request-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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

func acmResendValidationEmailCertificateCmd() *cobra.Command {
	var CertificateArn string
	var Domain string
	var ValidationDomain string

	cmd := &cobra.Command{
		Use: "resend-validation-email-certificates",
		Run: func(cmd *cobra.Command, args []string) {
			svc := acm.New(awsConfig)

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
