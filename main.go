package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

func main() {
	AwsCredentials, err := parseAwsCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(AwsCredentials)
}
