package credentials

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAwsCredentialsPath(t *testing.T) {
	testAwsCredentialsFile := filepath.Join("../", "test", ".aws", "credentials")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testAwsCredentialsFile)

	credentialsPath := getAwsCredentialsPath()
	if credentialsPath != testAwsCredentialsFile {
		t.Errorf("credentials file mismatch; actual %v, expected %v", credentialsPath, testAwsCredentialsFile)
	}
}

func TestGetAwsConfigPath(t *testing.T) {
	testAwsConfigFile := filepath.Join("../", "test", ".aws", "config")
	os.Setenv("AWS_CONFIG_FILE", testAwsConfigFile)

	configPath := getAwsConfigPath()
	if configPath != testAwsConfigFile {
		t.Errorf("config file mismatch; actual %v, expected %v", configPath, testAwsConfigFile)
	}
}

func TestParseAwsCredentialsAndConfig(t *testing.T) {
	testAwsCredentialsFile := filepath.Join("../", "test", ".aws", "credentials")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testAwsCredentialsFile)

	testAwsConfigFile := filepath.Join("../", "test", ".aws", "config")
	os.Setenv("AWS_CONFIG_FILE", testAwsConfigFile)

	ac, err := NewAwsCredentials()
	if err != nil {
		t.Error(err)
	}

	// Normal test
	if ac["default"].AccessKeyId != "AKIAIOSFODNN7EXAMPLE" {
		t.Errorf("access key id mismatch; actual %v, expected %v", ac["default"].AccessKeyId, "AKIAIOSFODNN7EXAMPLE")
	}

	if ac["default"].SecretAccessKey != "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" {
		t.Errorf("secret key mismatch; actual %v, expected %v", ac["default"].SecretAccessKey, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	}

	if ac["default"].Region != "us-east-1" {
		t.Errorf("region mismatch; actual %v, expected %v", ac["default"].Region, "us-east-1")
	}

	// Profile test
	if ac["for_profile_test"].AccessKeyId != "AKIAIOSFODNN8EXAMPLE" {
		t.Errorf("access key id mismatch; actual %v, expected %v", ac["for_profile_test"].AccessKeyId, "AKIAIOSFODNN8EXAMPLE")
	}

	if ac["for_profile_test"].SecretAccessKey != "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEZ" {
		t.Errorf("secret key mismatch; actual %v, expected %v", ac["for_profile_test"].SecretAccessKey, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEZ")
	}

	if ac["for_profile_test"].Region != "ap-northeast-1" {
		t.Errorf("region mismatch; actual %v, expected %v", ac["for_profile_test"].Region, "ap-northeast-1")
	}

	// Comment out test
	if _, ok := ac["for_comment_test"]; ok {
		t.Errorf("comment out test faild; profile %s", "for_comment_test")
	}

	// Only config test
	if _, ok := ac["for_null_test"]; ok {
		t.Errorf("only config test faild; profile %s", "for_null_test")
	}
}

func TestNotFoundAwsCredentials(t *testing.T) {
	testNotAwsCredentialsFile := filepath.Join("../", "test", ".aws", "not_credentials")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testNotAwsCredentialsFile)

	testAwsConfigFile := filepath.Join("../", "test", ".aws", "config")
	os.Setenv("AWS_CONFIG_FILE", testAwsConfigFile)

	if _, err := NewAwsCredentials(); err == nil {
		t.Error("credentials file should not be exists")
	}
}

func TestRegion(t *testing.T) {
	testRegionAwsCredentialsFile := filepath.Join("../", "test", ".aws", "credentials_region")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testRegionAwsCredentialsFile)

	ac, err := NewAwsCredentials()
	if err != nil {
		t.Error(err)
	}

	if ac["region_test"].AccessKeyId != "AKIAIOSFODNN0EXAMPLE" {
		t.Errorf("access key id mismatch; actual %v, expected %v", ac["region_test"].AccessKeyId, "AKIAIOSFODNN0EXAMPLE")
	}

	if ac["region_test"].SecretAccessKey != "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEB" {
		t.Errorf("secret key mismatch; actual %v, expected %v", ac["region_test"].SecretAccessKey, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEB")
	}

	if ac["region_test"].Region != "us-east-2" {
		t.Errorf("region mismatch; actual %v, expected %v", ac["region_test"].Region, "us-east-2")
	}
}
