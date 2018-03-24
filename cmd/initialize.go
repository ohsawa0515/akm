package cmd

import (
	"github.com/ohsawa0515/akm/config"
	"github.com/spf13/cobra"
)

func NewCmdInitialize() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize akm config file for the first time usage.",
		RunE:    initAction,
	}

	return cmd
}

func init() {}

func initAction(cmd *cobra.Command, args []string) error {
	if err := config.CreateAkmConfig(); err != nil {
		return err
	}

	cmd.Printf("akm config is created.\n")

	return nil
}
