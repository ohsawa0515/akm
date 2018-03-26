package cmd

import (
	"github.com/ohsawa0515/akm/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCmdCurrent() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "current",
		Aliases: []string{"c"},
		Short:   "Show current profile name.",
		RunE:    currentAction,
	}

	return cmd
}

func init() {}

func currentAction(cmd *cobra.Command, args []string) error {
	akmConfig, err := config.NewAkmConfig()
	if err != nil {
		return err
	}

	if len(akmConfig.Current) == 0 {
		return errors.Errorf("profile is not specified")
	}

	cmd.Printf(akmConfig.Current)

	return nil
}
