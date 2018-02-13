package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	shellwords "github.com/mattn/go-shellwords"
	cli "gopkg.in/urfave/cli.v1"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil || len(usr.HomeDir) > 0 {
		return usr.HomeDir
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("HOMEPATH")
	}
	return os.Getenv("HOME")
}

func getAwsCredentialsPath() string {
	if awsCredentialsFile, ok := os.LookupEnv("AWS_SHARED_CREDENTIALS_FILE"); ok {
		return awsCredentialsFile
	}

	return filepath.Join(getHomeDir(), ".aws", "credentials")
}

func getAwsConfigPath() string {
	if awsConfigFile, ok := os.LookupEnv("AWS_CONFIG_FILE"); ok {
		return awsConfigFile
	}

	return filepath.Join(getHomeDir(), ".aws", "config")
}

func initialize(c *cli.Context) error {
	if _, err := NewAkmConfig(); err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf(".akm.toml is created in %s\n", getAkmConfigPath())

	return nil
}

func list(c *cli.Context) error {
	ac, err := NewAwsCredentials(getAwsCredentialsPath(), getAwsConfigPath())
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	for profile := range ac {
		fmt.Println(profile)
	}

	return nil
}

func use(c *cli.Context) error {
	ac, err := NewAwsCredentials(getAwsCredentialsPath(), getAwsConfigPath())
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if c.NArg() == 0 {
		return cli.NewExitError("Select a profile", 2)
	}

	profile := c.Args().Get(0)
	if _, ok := ac[profile]; !ok {
		return cli.NewExitError(fmt.Sprintf("profile: %s doesn't exist", profile), 1)
	}

	if c.NArg() == 1 {
		var buf bytes.Buffer
		if len(ac[profile].AccessKeyId) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_ACCESS_KEY_ID='%s';", ac[profile].AccessKeyId))
		}
		if len(ac[profile].SecretAccessKey) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_SECRET_ACCESS_KEY='%s';", ac[profile].SecretAccessKey))
		}
		if len(ac[profile].Region) > 0 {
			buf.WriteString(fmt.Sprintf("export AWS_DEFAULT_REGION=%s", ac[profile].Region))
		}
		fmt.Println(buf.String())
	} else { // >= 2
		if len(ac[profile].AccessKeyId) > 0 {
			os.Setenv("AWS_ACCESS_KEY_ID", ac[profile].AccessKeyId)
		}
		if len(ac[profile].SecretAccessKey) > 0 {
			os.Setenv("AWS_SECRET_ACCESS_KEY", ac[profile].SecretAccessKey)
		}
		if len(ac[profile].Region) > 0 {
			os.Setenv("AWS_DEFAULT_REGION", ac[profile].Region)
		}

		var buf bytes.Buffer
		for i := 1; i < c.NArg(); i++ {
			buf.WriteString(c.Args().Get(i))
			buf.WriteString(" ")
		}
		command := buf.String()

		args, err := shellwords.Parse(command)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Println(string(out))
	}

	return nil
}

func configure(c *cli.Context) error {
	ac, err := NewAwsCredentials(getAwsCredentialsPath(), getAwsConfigPath())
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if c.NArg() == 0 {
		return cli.NewExitError("Specify a profile", 2)
	}

	profile := c.Args().Get(0)
	if _, ok := ac[profile]; !ok {
		ac[profile] = &AwsCredential{}
	}

	// Start prompt
	if err := ac[profile].AccessKeyIdPrompt(); err != nil {
		return cli.NewExitError(err, 1)
	}
	if err := ac[profile].SecretAccessKeyPrompt(); err != nil {
		return cli.NewExitError(err, 1)
	}
	if err := ac[profile].RegionPrompt(); err != nil {
		return cli.NewExitError(err, 1)
	}

	// Save to credentials file
	if err := ac[profile].SaveToCredentialsFilePrompt(profile, getAwsCredentialsPath()); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func delete(c *cli.Context) error {
	ac, err := NewAwsCredentials(getAwsCredentialsPath(), getAwsConfigPath())
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if c.NArg() == 0 {
		return cli.NewExitError("Specify a profile", 2)
	}

	profile := c.Args().Get(0)
	if _, ok := ac[profile]; !ok {
		return cli.NewExitError(fmt.Sprintf("profile: %s doesn't exist", profile), 1)
	}

	// Delete profile from credentials file
	if err := ac[profile].DeleteProfilePrompt(profile, getAwsCredentialsPath()); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func clear(c *cli.Context) error {
	var buf bytes.Buffer
	buf.WriteString("unset AWS_ACCESS_KEY_ID;")
	buf.WriteString("unset AWS_SECRET_ACCESS_KEY;")
	buf.WriteString("unset AWS_DEFAULT_REGION;")
	fmt.Println(buf.String())

	return nil
}
