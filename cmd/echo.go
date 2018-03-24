package cmd

import (
	"github.com/ohsawa0515/akm/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/urfave/cli"
)

func NewCmdEcho() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "echo",
		Aliases: []string{"e"},
		Short:   "Show the AWS key or region with the specified profile name",
		Args:    cobra.MinimumNArgs(2),
		RunE:    echoAction,
	}

	return cmd
}

func init() {}

func echoAction(cmd *cobra.Command, args []string) error {
	ac, err := credentials.NewAwsCredentials()
	if err != nil {
		return err
	}

	switch len(args) {
	case 0:
		return cli.NewExitError("Select a profile", 2)
	case 1:
		return cli.NewExitError("Select a setting (aws_access_key_id or aws_secret_access_key or region)", 2)
	}

	profile := args[0]
	if _, ok := ac[profile]; !ok {
		return errors.Errorf("profile: %s doesn't exist", profile)
	}

	setting := args[1]
	switch setting {
	case "aws_access_key_id":
		cmd.Println(ac[profile].AccessKeyId)
		return nil
	case "aws_secret_access_key":
		cmd.Println(ac[profile].SecretAccessKey)
		return nil
	case "region":
		cmd.Println(ac[profile].Region)
		return nil
	default:
		return errors.Errorf("Select a setting (aws_access_key_id or aws_secret_access_key or region)")
	}

	return nil
}
