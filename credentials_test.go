package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseAwsCredentialsAndConfig(t *testing.T) {
	dir, _ := os.Getwd()
	ac, err := NewAwsCredentials(filepath.Join(dir, "test", ".aws", "credentials"), filepath.Join(dir, "test", ".aws", "config"))
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
	dir, _ := os.Getwd()
	if _, err := NewAwsCredentials(filepath.Join(dir, "test", ".aws", "not_credentials"), filepath.Join(dir, "test", ".aws", "config")); err == nil {
		t.Error("credentials file should not be exists")
	}

	if _, err := NewAwsCredentials(filepath.Join(dir, "test", ".aws", "credentials"), filepath.Join(dir, "test", ".aws", "not_config")); err == nil {
		t.Error("config file should not be exists")
	}
}
