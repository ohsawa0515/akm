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
		Short:   "Delete the environment variable of AWS credentials.",
		Example: `$ akm clear
  unset AWS_ACCESS_KEY_ID;unset AWS_SECRET_ACCESS_KEY;unset AWS_DEFAULT_REGION;

  Delete environment variable with eval.
  $ env | grep AWS
  AWS_ACCESS_KEY_ID=xxxxxxx
  AWS_SECRET_ACCESS_KEY=xxxxxxx
  AWS_DEFAULT_REGION=us-east-1

  $ eval $(akm clear)

  $ env | grep AWS
  # empty`,
		RunE: clearAction,
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
