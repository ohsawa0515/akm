package cmd

import (
	"os"

	"github.com/ohsawa0515/akm/app"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var appName = "akm"

func NewCmdRoot() *cobra.Command {
	cliApp := app.NewCliApps()

	cmd := &cobra.Command{
		Use:              appName + " COMMAND",
		Short:            "A simple AWS access keys manager",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		Args:             cobra.NoArgs,
		Run:              func(cmd *cobra.Command, args []string) {},
	}

	addCommands(cmd)

	cobra.AddTemplateFunc("useLine", UseLine)
	cobra.AddTemplateFunc("version", func() string { return cliApp.Version })
	cobra.AddTemplateFunc("author", func() string { return cliApp.Author })
	cmd.SetUsageTemplate(usageTemplate)
	cmd.SetHelpTemplate(helpTemplate)

	cmd.PersistentFlags().BoolP("help", "h", false, "Print help")

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
	cmd.AddCommand(NewCmdCurrent())
	cmd.AddCommand(NewCmdEcho())
	cmd.AddCommand(NewCmdConfigure())
	cmd.AddCommand(NewCmdDelete())
	cmd.AddCommand(NewCmdClear())

	var useSub bool
	if len(os.Args) > 3 {
		useSub = true
	} else {
		useSub = false
	}
	cmd.AddCommand(NewCmdUse(useSub))
}

func UseLine(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return cmd.Parent().CommandPath() + " " + cmd.Use
	}
	return appName + " " + cmd.Use
}

var usageTemplate = `
Usage:
{{- if not .HasSubCommands}}
  {{ useLine . }}
{{- end}}
{{- if .HasSubCommands}}
  {{ .CommandPath}} COMMAND

Version:
  {{ version }}

Author:
  {{ author}}
{{- end}}
{{- if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{- end}}
{{- if gt (len .Long) 0}}

Description:
  {{ .Long }}
{{- end}}
{{- if .HasExample}}

Examples:
  {{ .Example }}
{{- end}}
{{- if gt (len .LocalFlags.FlagUsages) 0}}

Commands:
{{- range .Commands}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{- end}}
{{- if gt (len .InheritedFlags.FlagUsages) 0}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{- end}}
`

var helpTemplate = `Name:
  {{ .CommandPath}} - {{ .Short | trim }}
{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`
