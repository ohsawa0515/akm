package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mattn/go-shellwords"
)

func setUpClear() {
	testAkmConfigPath := filepath.Join("../", "test", ".akm", "config_clear")

	os.Setenv("AKM_CONFIG_DIR", filepath.Join("../", "test", ".akm"))
	os.Setenv("AKM_CONFIG_FILE", testAkmConfigPath)

	f, _ := os.OpenFile(testAkmConfigPath, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
}

func TestCmdClear(t *testing.T) {
	setUpClear()

	cases := []struct {
		command string
	}{
		{command: "akm use for_profile_test"},
		{command: "akm clear"},
	}
	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
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
