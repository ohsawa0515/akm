package app

type cliApp struct {
	Version string
	Revision string
	Author  string
}

func NewCliApps() cliApp {
	return cliApp{
		Version: Version,
		Revision: Revision,
		Author:  "Shuichi Ohsawa <ohsawa0515@gmail.com>",
	}
}
