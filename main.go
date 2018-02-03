package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-ini/ini"
	shellwords "github.com/mattn/go-shellwords"
	cli "gopkg.in/urfave/cli.v1"
)

type AwsCredential struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}

type AwsCredentials map[string]*AwsCredential

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil || len(usr.HomeDir) > 0 {
		return usr.HomeDir
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("HOMEPATH")
	}
	return os.Getenv("HOME")
}

func getAwsCredentialsPath() string {
	return filepath.Join(getHomeDir(), ".aws", "credentials")
}

func getAwsConfigPath() string {
	return filepath.Join(getHomeDir(), ".aws", "config")
}

func parseAwsCredentials() (AwsCredentials, error) {

	awsCredentials := make(AwsCredentials)

	// Credentials file
	cre, err := ini.Load(getAwsCredentialsPath())
	if err != nil {
		return nil, err
	}
	cre.BlockMode = false

	for _, section := range cre.Sections() {
		awsCre := &AwsCredential{}
		profile := section.Name()
		if profile == "DEFAULT" {
			continue
		}
		awsAccessKeyId, err := section.GetKey("aws_access_key_id")
		if err == nil {
			awsCre.AccessKeyId = awsAccessKeyId.String()
		}

		awsSecretAccessKey, err := section.GetKey("aws_secret_access_key")
		if err == nil {
			awsCre.SecretAccessKey = awsSecretAccessKey.String()
		}
		awsCredentials[profile] = awsCre
	}

	// Config file
	config, err := ini.Load(getAwsConfigPath())
	if err != nil {
		return nil, err
	}
	config.BlockMode = false

	for _, section := range config.Sections() {
		profile := strings.Replace(section.Name(), "profile ", "", -1)
		if profile == "DEFAULT" {
			continue
		}
		if _, ok := awsCredentials[profile]; !ok {
			continue
		}
		region, err := section.GetKey("region")
		if err == nil {
			awsCredentials[profile].Region = region.String()
		}
	}

	return awsCredentials, nil
}

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

func main() {
	app := cli.NewApp()
	app.Name = "akm"
	app.Usage = "A simple AWS access keys manager"
	app.Author = "Shuichi Ohsawa"
	app.Email = "ohsawa0515@gmail.com"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "List all AWS credentials profile",
			Action:  list,
		},
		{
			Name:    "use",
			Aliases: []string{"u"},
			Usage:   "Set specific AWS credential in environment values",
			Action:  use,
		},
	}
	app.Run(os.Args)
}
