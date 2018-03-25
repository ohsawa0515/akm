package app

type cliApp struct {
	Version string
	Author  string
}

func NewCliApps() cliApp {
	return cliApp{
		Version: version,
		Author:  "Shuichi Ohsawa <ohsawa0515@gmail.com>",
	}
}
