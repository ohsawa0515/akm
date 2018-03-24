package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "akm",
		Short: "A simple AWS access keys manager",

		Run: func(cmd *cobra.Command, args []string) {},
	}
	addCommands(cmd)

	return cmd
}

func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func addCommands(cmd *cobra.Command) {
	cmd.AddCommand(NewCmdInitialize())
	cmd.AddCommand(NewCmdList())
	cmd.AddCommand(NewCmdUse())
	cmd.AddCommand(NewCmdCurrent())
	cmd.AddCommand(NewCmdEcho())
	cmd.AddCommand(NewCmdConfigure())
	cmd.AddCommand(NewCmdDelete())
	cmd.AddCommand(NewCmdClear())
}
