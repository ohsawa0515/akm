package cmd

import (
	"bytes"

	"github.com/ohsawa0515/akm/config"
	"github.com/spf13/cobra"
)

func NewCmdClear() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clear",
		Aliases: []string{"C"},
		Short:   "Clear the environment variable of AWS credentials",
		RunE:    clearAction,
	}

	return cmd
}

func clearAction(cmd *cobra.Command, args []string) error {
	var buf bytes.Buffer
	buf.WriteString("unset AWS_ACCESS_KEY_ID;")
	buf.WriteString("unset AWS_SECRET_ACCESS_KEY;")
	buf.WriteString("unset AWS_DEFAULT_REGION;")
	cmd.Println(buf.String())

	akmConfig, err := config.NewAkmConfig()
	if err != nil {
		return err
	}
	if err := akmConfig.Delete(); err != nil {
		return err
	}

	return nil
}
