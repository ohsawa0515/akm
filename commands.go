package main

import cli "gopkg.in/urfave/cli.v1"

func commands() []cli.Command {
	return []cli.Command{
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
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   "Clear the environment variable of AWS credentials",
			Action:  clear,
		},
	}
}
