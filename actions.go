package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	shellwords "github.com/mattn/go-shellwords"
	cli "gopkg.in/urfave/cli.v1"
)

func list(c *cli.Context) error {
	AwsCredentials, err := parseAwsCredentials()
	if err != nil {
		return err
	}

	if len(AwsCredentials) == 0 {
		fmt.Println("No AWS credentials found.")
		return nil
	}

	for profile := range AwsCredentials {
		fmt.Println(profile)
	}

	return nil
}

func use(c *cli.Context) error {
	awsCredentials, err := parseAwsCredentials()
	if err != nil {
		return err
	}

	if len(awsCredentials) == 0 {
		fmt.Println("No AWS credentials found.")
		return nil
	}

	if c.NArg() == 0 {
		fmt.Println("Select a profile")
		return nil
	}

	profile := c.Args().Get(0)
	_, ok := awsCredentials[profile]
	if !ok {
		return fmt.Errorf("Profile: %s doesn't exists\n", profile)
	}

	if c.NArg() == 1 {
		var buf bytes.Buffer
		if len(awsCredentials[profile].AccessKeyId) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_ACCESS_KEY_ID='%s';", awsCredentials[profile].AccessKeyId))
		}
		if len(awsCredentials[profile].SecretAccessKey) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_SECRET_ACCESS_KEY='%s';", awsCredentials[profile].SecretAccessKey))
		}
		if len(awsCredentials[profile].Region) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_DEFAULT_REGION=%s", awsCredentials[profile].Region))
		}
		fmt.Println(buf.String())
	} else { // >= 2
		if len(awsCredentials[profile].AccessKeyId) > 0 {
			os.Setenv("AWS_ACCESS_KEY_ID", awsCredentials[profile].AccessKeyId)
		}
		if len(awsCredentials[profile].SecretAccessKey) > 0 {
			os.Setenv("AWS_SECRET_ACCESS_KEY", awsCredentials[profile].SecretAccessKey)
		}
		if len(awsCredentials[profile].Region) > 0 {
			os.Setenv("AWS_DEFAULT_REGION", awsCredentials[profile].Region)
		}

		var buf bytes.Buffer
		for i := 1; i < c.NArg(); i++ {
			buf.WriteString(c.Args().Get(i))
			buf.WriteString(" ")
		}
		command := buf.String()

		args, err := shellwords.Parse(command)
		if err != nil {
			return err
		}

		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}

	return nil
}
