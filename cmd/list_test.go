package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestCmdList(t *testing.T) {
	setUp()

	cases := []struct {
		command string
	}{
		{command: "akm ls"},
		{command: "akm l"},
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
