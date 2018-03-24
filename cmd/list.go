package cmd

import (
	"sort"

	"github.com/ohsawa0515/akm/config"
	"github.com/ohsawa0515/akm/credentials"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "List all AWS credentials profile",
		RunE:    listAction,
	}

	return cmd
}

func init() {}

func listAction(cmd *cobra.Command, args []string) error {
	ac, err := credentials.NewAwsCredentials()
	if err != nil {
		return err
	}

	akmConfig, err := config.NewAkmConfig()
	if err != nil {
		return err
	}

	var profiles []string
	for p := range ac {
		profiles = append(profiles, p)
	}
	sort.Strings(profiles)

	for _, profile := range profiles {
		if profile == akmConfig.Current {
			cmd.Printf("%s (Current)\n", profile)
		} else {
			cmd.Printf("%s\n", profile)
		}
	}

	return nil
}
