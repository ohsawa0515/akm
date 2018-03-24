package config

import (
	"os"
	"path/filepath"
	"testing"
)

func getTestAkmHomeDir() string {
	d := filepath.Join("../", "test", ".akm")
	os.Setenv("AKM_CONFIG_DIR", d)

	return d
}

func getTestAkmConfigPath() string {
	d := filepath.Join("../", "test", ".akm", "config")
	os.Setenv("AKM_CONFIG_FILE", d)

	return d
}

func TestAkmHomeDir(t *testing.T) {
	testAkmHomeDir := getTestAkmHomeDir()
	akmHomeDir := getAkmHomeDir()
	if akmHomeDir != testAkmHomeDir {
		t.Errorf("akm honme dir path mismatch; actual %v, expected %v", akmHomeDir, testAkmHomeDir)
	}
}

func TestAkmConfigPath(t *testing.T) {
	testAkmConfigPath := getTestAkmConfigPath()
	akmConfigPath := getAkmConfigPath()
	if akmConfigPath != testAkmConfigPath {
		t.Errorf("akm config file path mismatch; actual %v, expected %v", akmConfigPath, testAkmConfigPath)
	}
}

func TestAkmConfigAlreadyExists(t *testing.T) {
	getTestAkmHomeDir()    // Read Environment variables
	getTestAkmConfigPath() // Read Environment variables
	if err := CreateAkmConfig(); err == nil {
		t.Errorf("expected akm config file is already exists.")
	}
}

func TestDecodeAkmConfig(t *testing.T) {
	getTestAkmHomeDir()    // Read Environment variables
	getTestAkmConfigPath() // Read Environment variables
	akmConfig, err := NewAkmConfig()
	if err != nil {
		t.Error(err)
	}
	if akmConfig.Current != "for_profile_test" {
		t.Errorf("akm config file decode mismatch; actual %v, expected %v", akmConfig.Current, "for_profile_test")
	}
}
