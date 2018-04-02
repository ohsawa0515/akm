package credentials

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/ohsawa0515/akm/utils"
	"github.com/pkg/errors"
)

type AwsCredential struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}

type AwsCredentials map[string]*AwsCredential

func getAwsCredentialsPath() string {
	if awsCredentialsFile, ok := os.LookupEnv("AWS_SHARED_CREDENTIALS_FILE"); ok {
		return awsCredentialsFile
	}

	return filepath.Join(utils.GetHomeDir(), ".aws", "credentials")
}

func getAwsConfigPath() string {
	if awsConfigFile, ok := os.LookupEnv("AWS_CONFIG_FILE"); ok {
		return awsConfigFile
	}

	return filepath.Join(utils.GetHomeDir(), ".aws", "config")
}

func NewAwsCredentials() (AwsCredentials, error) {
	ac := make(AwsCredentials)

	acPath := getAwsCredentialsPath()
	if !utils.IsExist(acPath) {
		return nil, errors.Errorf("credentials file not found")
	}
	if err := ac.ParseAwsCredentials(acPath); err != nil {
		return nil, err
	}

	conPath := getAwsConfigPath()
	if utils.IsExist(conPath) {
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
