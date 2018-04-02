package cmd

import (
	"github.com/ohsawa0515/akm/config"
	"github.com/spf13/cobra"
)

func NewCmdInitialize() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize for akm command",
		Long: `Initialize akm command for the first time usage.
  After execution, "$HOME/.akm.toml" is created.`,
		RunE: initAction,
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
