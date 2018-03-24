package cmd

import (
	"github.com/ohsawa0515/akm/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "d"},
		Short:   "Delete profile from AWS credentials file",
		Args:    cobra.MinimumNArgs(1),
		RunE:    deleteAction,
	}

	return cmd
}

func deleteAction(cmd *cobra.Command, args []string) error {
	ac, err := credentials.NewAwsCredentials()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.Errorf("Specify a profile")
	}

	profile := args[0]
	if _, ok := ac[profile]; !ok {
		return errors.Errorf("profile: %s doesn't exist", profile)
	}

	// Delete profile from credentials file
	if err := ac[profile].DeleteProfilePrompt(profile); err != nil {
		return err
	}

	return nil
}
