package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestCmdEcho(t *testing.T) {
	setUp()

	cases := []struct {
		command string
		want    string
	}{
		{command: "akm echo for_profile_test aws_access_key_id", want: "AKIAIOSFODNN8EXAMPLE\n"},
		{command: "akm echo for_profile_test aws_secret_access_key", want: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEZ\n"},
		{command: "akm echo for_profile_test region", want: "ap-northeast-1\n"},
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
		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
