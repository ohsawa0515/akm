package main

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-ini/ini"
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
