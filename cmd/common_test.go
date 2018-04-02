package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func setUp() {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", filepath.Join("../", "test", ".aws", "credentials"))
	os.Setenv("AWS_CONFIG_FILE", filepath.Join("../", "test", ".aws", "config"))
	os.Setenv("AKM_CONFIG_FILE", filepath.Join("../", "test", ".akm", "config"))
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
