package cmd

import (
	"github.com/ohsawa0515/akm/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCmdConfigure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure AWS credentials",
		Args:  cobra.MinimumNArgs(1),
		RunE:  configureAction,
	}

	return cmd
}

func configureAction(cmd *cobra.Command, args []string) error {
	ac, err := credentials.NewAwsCredentials()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.Errorf("Specify a profile")
	}

	profile := args[0]
	if _, ok := ac[profile]; !ok {
		ac[profile] = &credentials.AwsCredential{}
	}

	// Start prompt
	if err := ac[profile].AccessKeyIdPrompt(); err != nil {
		return err
	}
	if err := ac[profile].SecretAccessKeyPrompt(); err != nil {
		return err
	}
	if err := ac[profile].RegionPrompt(); err != nil {
		return err
	}

	// Save to credentials file
	if err := ac[profile].SaveToCredentialsFilePrompt(profile); err != nil {
		return err
	}

	return nil
}
