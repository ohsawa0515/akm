package main

import cli "gopkg.in/urfave/cli.v1"

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize akm config file for the first time usage",
			Action:  initialize,
		},
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "List all AWS credentials profile",
			Action:  list,
		},
		{
			Name:    "use",
			Aliases: []string{"u"},
			Usage:   "Set specific AWS credential in environment values",
			Action:  use,
		},
		{
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "Show current profile name",
			Action:  current,
		},
		{
			Name:    "echo",
			Aliases: []string{"e"},
			Usage:   "Show the current AWS key or region by profile name",
			Action:  echo,
		},
		{
			Name:   "clear",
			Usage:  "Clear the environment variable of AWS credentials",
			Action: clear,
		},
		{
			Name:    "configure",
			Aliases: []string{"c"},
			Usage:   "Configure AWS credentials",
			Action:  configure,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete profile from AWS credentials file",
			Action:  delete,
		},
	}
}
