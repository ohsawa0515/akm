package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mattn/go-shellwords"
)

func setUpUse() {
	testAkmConfigPath := filepath.Join("../", "test", ".akm", "config_use")

	os.Setenv("AKM_CONFIG_DIR", filepath.Join("../", "test", ".akm"))
	os.Setenv("AKM_CONFIG_FILE", testAkmConfigPath)

	f, _ := os.OpenFile(testAkmConfigPath, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
}

func TestCmdUseSuccess(t *testing.T) {
	setUpUse()

	cases := []struct {
		command string
	}{
		{command: "akm use for_profile_test ls"},
		{command: "akm use for_profile_test ls -la"},
		{command: "akm use for_profile_test cd ../"},
		{command: "akm use for_profile_test alias ll='ls -lG'"},
	}
	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()

		// Only use command
		for _, cmd := range cmd.Commands() {
			cmd.DisableFlagParsing = true
		}

		cmd.SetOutput(buf)
		cmdArgs, err := shellwords.Parse(c.command)
		if err != nil {
			t.Errorf("args parse error: %+v\n", err)
		}
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		if err := cmd.Execute(); err != nil {
			t.Errorf("unexpected error:%+v", err)
		}
	}
}

func TestCmdUseFailure(t *testing.T) {
	setUpUse()

	cases := []struct {
		command string
	}{
		{command: "akm use"},                              // Require arguments
		{command: "akm use foo"},                          // Missing profile
		{command: "akm use for_profile_test foo_command"}, // Invalid command
	}
	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()

		// Only use command
		for _, cmd := range cmd.Commands() {
			cmd.DisableFlagParsing = true
		}

		cmd.SetOutput(buf)
		cmdArgs, err := shellwords.Parse(c.command)
		if err != nil {
			t.Errorf("args parse error: %+v\n", err)
		}
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		if err := cmd.Execute(); err == nil {
			t.Errorf("unexpected not error:%+v", err)
		}
	}
}
