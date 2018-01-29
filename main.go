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

var AwsCredentials []AwsCredential

const Secret = "****"

type AwsCredential struct {
	Profile         string
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

func parseAwsCredentials() ([]AwsCredential, error) {
	awsConfigurePath := filepath.Join(getHomeDir(), ".aws")
	awsCredentialsPath := filepath.Join(awsConfigurePath, "credentials")
	f, err := os.Open(awsCredentialsPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reAccessKeyId := regexp.MustCompile(`aws_access_key_id\s*?=\s*?(\w.*)`)
	reSecretAccessKey := regexp.MustCompile(`aws_secret_access_key\s*?=\s*?(\w.*)`)
	var awsCres []AwsCredential

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
			awsCre := AwsCredential{
				Profile: profile,
			}

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

			awsCres = append(awsCres, awsCre)
		}
	}

	return awsCres, nil
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

	for _, awsCre := range AwsCredentials {
		fmt.Println(awsCre.Profile)
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
			Usage:   "List all AWS credentials profile.",
			Action:  list,
		},
	}
	app.Run(os.Args)
}
