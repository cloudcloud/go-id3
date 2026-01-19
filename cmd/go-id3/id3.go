// Package main provides a library for working with ID3 tags in Audio and Video files.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudcloud/go-id3/internal/version"
	"github.com/urfave/cli/v3"
)

const (
	appHelpTemplate = `Usage:
  {{.Name}} <command> [options...]

Available commands are: {{range .VisibleCategories}}{{if .Name}}
{{.Name}}:{{range .VisibleCommands}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{"\n"}}{{else}}{{range .VisibleCommands}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{"\n"}}{{end}}{{end}}
Use "{{.Name}} <command> --help" for more information about a command.
`

	commandHelpTemplate = `{{.Description}}

Options:

{{range .VisibleFlags}}  {{.}}
{{ end -}}
`
)

func printVersion(cmd *cli.Command) {
	fmt.Fprintf(cmd.Root().Writer, "%s: %s\n", cmd.Root().Name, version.FullVersion())
}

func init() {
	cli.RootCommandHelpTemplate = appHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate

	cli.VersionFlag = &cli.BoolFlag{Name: "print-version", Aliases: []string{"v"}}

	cli.VersionPrinter = printVersion
	cli.ErrWriter = os.Stderr
}

func main() {
	cmd := &cli.Command{
		Name:            "go-id3",
		Version:         version.FullVersion(),
		HideVersion:     false,
		SkipFlagParsing: false,
		Commands: []*cli.Command{
			readCmd,
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "go-id3 fatal: %s\n", err)
	}
}
