package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

type AwsCredential struct {
	AccessKeyId     string
	SecretAccessKey string
}

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

func parseAwsCredentials() (map[string]*AwsCredential, error) {
	awsConfigurePath := filepath.Join(getHomeDir(), ".aws")
	awsCredentialsPath := filepath.Join(awsConfigurePath, "credentials")
	f, err := os.Open(awsCredentialsPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reAccessKeyId := regexp.MustCompile(`aws_access_key_id\s*?=\s*?(\w.*)`)
	reSecretAccessKey := regexp.MustCompile(`aws_secret_access_key\s*?=\s*?(\w.*)`)
	awsCredentials := make(map[string]*AwsCredential)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip blank line
		if len(line) == 0 {
			continue
		}

		// Skip comment out
		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			tmp := strings.Trim(line, "[")
			profile := strings.Trim(tmp, "]")
			awsCre := &AwsCredential{}

			scanner.Scan()
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "aws_access_key_id") {
				matches := reAccessKeyId.FindStringSubmatch(line)
				awsCre.AccessKeyId = matches[1]
			}

			scanner.Scan()
			line = strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "aws_secret_access_key") {
				matches := reSecretAccessKey.FindStringSubmatch(line)
				awsCre.SecretAccessKey = matches[1]
			}

			awsCredentials[profile] = awsCre

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

	if c.NArg() > 0 {
		profile := c.Args().Get(0)

		_, ok := awsCredentials[profile]
		if !ok {
			fmt.Printf("Profile: %s doesn't exists\n", profile)
			return nil
		}

		out := ""
		out += fmt.Sprintf("export AWS_ACCESS_KEY_ID='%s';", awsCredentials[profile].AccessKeyId)
		out += fmt.Sprintf("export AWS_SECRET_ACCESS_KEY='%s';", awsCredentials[profile].SecretAccessKey)
		fmt.Println(out)

	} else {
		fmt.Println("Select a profile")
		return nil
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
