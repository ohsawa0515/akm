package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAwsCredentialsPath(t *testing.T) {
	dir, _ := os.Getwd()
	testAwsCredentialsFile := filepath.Join(dir, "test", ".aws", "credentials")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testAwsCredentialsFile)

	credentialsPath := getAwsCredentialsPath()
	if credentialsPath != testAwsCredentialsFile {
		t.Errorf("credentials file mismatch; actual %v, expected %v", credentialsPath, testAwsCredentialsFile)
	}
}

func TestGetAwsConfigPath(t *testing.T) {
	dir, _ := os.Getwd()
	testAwsConfigFile := filepath.Join(dir, "test", ".aws", "config")
	os.Setenv("AWS_CONFIG_FILE", testAwsConfigFile)

	configPath := getAwsConfigPath()
	if configPath != testAwsConfigFile {
		t.Errorf("config file mismatch; actual %v, expected %v", configPath, testAwsConfigFile)
	}
}
