package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/ohsawa0515/akm/config"
	"github.com/ohsawa0515/akm/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCmdUse() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "use",
		Aliases:            []string{"u"},
		Short:              "Set specific AWS credential in environment values",
		DisableFlagParsing: true,
		Args:               cobra.MinimumNArgs(1),
		RunE:               useAction,
	}

	return cmd
}

func init() {}

func useAction(cmd *cobra.Command, args []string) error {
	ac, err := credentials.NewAwsCredentials()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.Errorf("Select a profile")
	}

	profile := args[0]
	if _, ok := ac[profile]; !ok {
		return errors.Errorf("profile: %s doesn't exist", profile)
	}

	if len(args) == 1 { // Set environment variables
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
		cmd.Println(buf.String())
	} else { // >= 2  execute any command
		if len(ac[profile].AccessKeyId) > 0 {
			os.Setenv("AWS_ACCESS_KEY_ID", ac[profile].AccessKeyId)
		}
		if len(ac[profile].SecretAccessKey) > 0 {
			os.Setenv("AWS_SECRET_ACCESS_KEY", ac[profile].SecretAccessKey)
		}
		if len(ac[profile].Region) > 0 {
			os.Setenv("AWS_DEFAULT_REGION", ac[profile].Region)
		}

		out, err := exec.Command(args[1], args[2:]...).CombinedOutput()
		if err != nil {
			return err
		}
		cmd.Println(string(out))
	}

	// Set current setting to config file
	akmConfig, err := config.NewAkmConfig()
	if err != nil {
		return err
	}
	akmConfig.Current = profile
	if err := akmConfig.Save(); err != nil {
		return err
	}

	return nil
}
