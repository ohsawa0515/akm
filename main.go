package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "akm"
	app.Usage = "A simple AWS access keys manager"
	app.Author = "Shuichi Ohsawa"
	app.Email = "ohsawa0515@gmail.com"
	app.Version = "0.1.0"
	app.Commands = commands()
	app.Run(os.Args)
}
