package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

type AwsCredential struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}

type AwsCredentials map[string]*AwsCredential

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func NewAwsCredentials(acPath, conPath string) (AwsCredentials, error) {
	ac := make(AwsCredentials)

	if err := ac.ParseAwsCredentials(acPath); err != nil {
		return nil, err
	}

	if isExist(conPath) {
		if err := ac.ParseAwsConfig(conPath); err != nil {
			return nil, err
		}
	}

	if len(ac) == 0 {
		return nil, fmt.Errorf("no AWS credentials found")
	}

	return ac, nil
}

func (ac AwsCredentials) ParseAwsCredentials(acPath string) error {

	cre, err := ini.Load(acPath)
	if err != nil {
		return err
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

		awsRegion, err := section.GetKey("region")
		if err == nil {
			awsCre.Region = awsRegion.String()
		}
		ac[profile] = awsCre
	}

	return nil
}

func (ac AwsCredentials) ParseAwsConfig(conPath string) error {

	config, err := ini.Load(conPath)
	if err != nil {
		return err
	}
	config.BlockMode = false

	for _, section := range config.Sections() {
		profile := strings.Replace(section.Name(), "profile ", "", -1)
		if profile == "DEFAULT" {
			continue
		}
		if _, ok := ac[profile]; !ok {
			continue
		}
		region, err := section.GetKey("region")
		if err == nil {
			ac[profile].Region = region.String()
		}
	}

	return nil
}
