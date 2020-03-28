package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/applicreation/aws-cli/settings"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Config struct {
	Profile         string
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}

func configureCmd(options *settings.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure aws-cli",
		Run: func(cmd *cobra.Command, args []string) {
			config := Config{
				Profile: options.Profile,
			}

			cfg, err := external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(options.Profile))
			if err == nil {
				ctx := context.Background()
				credentials, _ := cfg.Credentials.Retrieve(ctx)

				config.AccessKeyId = credentials.AccessKeyID
				config.SecretAccessKey = credentials.SecretAccessKey
				config.Region = cfg.Region
			}

			config.AccessKeyId, err = ReadStringFromUser("AWS Access Key ID", true, config.AccessKeyId)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			config.SecretAccessKey, err = ReadStringFromUser("AWS Secret Access Key", true, config.SecretAccessKey)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			config.Region, err = ReadStringFromUser("Default region name", false, config.Region)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			err = WriteCredentials(config)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}

			err = WriteConfig(config)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func ReadStringFromUser(message string, mask bool, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label: message,
	}

	if mask == true {
		prompt.Mask = '*'
	}

	if defaultValue != "" {
		prompt.Default = defaultValue
	}

	value, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return value, nil
}

func WriteCredentials(config Config) error {
	file, err := os.OpenFile(external.DefaultSharedCredentialsFilename(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}

	lineNumberStart := FindSectionStart(data, config.Profile)
	lineNumberEnd := FindSectionEnd(data, lineNumberStart)
	if lineNumberStart >= 0 && lineNumberEnd >= 0 {
		UpdateCredentials(file, data, lineNumberStart, lineNumberEnd, config)
	} else {
		WriteNewCredentials(file, data, config)
	}

	return nil
}

func WriteConfig(config Config) error {
	file, err := os.OpenFile(external.DefaultSharedConfigFilename(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}

	lineNumberStart := FindSectionStart(data, config.Profile)
	lineNumberEnd := FindSectionEnd(data, lineNumberStart)
	if lineNumberStart >= 0 && lineNumberEnd >= 0 {
		UpdateConfig(file, data, lineNumberStart, lineNumberEnd, config)
	} else {
		WriteNewConfig(file, data, config)
	}

	return nil
}

func FindSectionStart(data []byte, profile string) int {
	line := 0

	file := string(data)

	temp := strings.Split(file, "\n")

	for _, item := range temp {
		if strings.Contains(item, profile) {
			return line
		}

		line++
	}

	return -1
}

func FindSectionEnd(data []byte, lineNumberStart int) int {
	if lineNumberStart < 0 {
		return lineNumberStart
	}

	line := 0

	file := string(data)

	temp := strings.Split(file, "\n")

	for _, item := range temp {
		if line > lineNumberStart && item == "" {
			return line
		}

		line++
	}

	return line
}

func WriteNewCredentials(file *os.File, data []byte, config Config) {
	lines := len(strings.Split(strings.TrimSpace(string(data)), "\n"))

	if lines > 1 {
		file.WriteString("\n")
	} else {
		file.Truncate(0)
	}

	if _, err := file.WriteString("[" + config.Profile + "]\n"); err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString("aws_access_key_id = " + config.AccessKeyId + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString("aws_secret_access_key = " + config.SecretAccessKey + "\n"); err != nil {
		log.Println(err)
	}
}

func UpdateCredentials(file *os.File, data []byte, lineNumberStart int, lineNumberEnd int, config Config) {
	currentLine := 0

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if currentLine >= lineNumberStart && currentLine <= lineNumberEnd {
			if strings.Contains(line, "aws_access_key_id") {
				lines[i] = "aws_access_key_id = " + config.AccessKeyId
			}
			if strings.Contains(line, "aws_secret_access_key") {
				lines[i] = "aws_secret_access_key = " + config.SecretAccessKey
			}
		}

		currentLine++
	}

	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file.Name(), []byte(output), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteNewConfig(file *os.File, data []byte, config Config) {
	lines := len(strings.Split(strings.TrimSpace(string(data)), "\n"))

	if lines > 1 {
		file.WriteString("\n")
	} else {
		file.Truncate(0)
	}

	if _, err := file.WriteString("[" + config.Profile + "]\n"); err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString("output = json\n"); err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString("region = " + config.Region + "\n"); err != nil {
		log.Println(err)
	}
}

func UpdateConfig(file *os.File, data []byte, lineNumberStart int, lineNumberEnd int, config Config) {
	currentLine := 0

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if currentLine >= lineNumberStart && currentLine <= lineNumberEnd {
			if strings.Contains(line, "region") {
				lines[i] = "region = " + config.Region
			}
		}

		currentLine++
	}

	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file.Name(), []byte(output), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}
