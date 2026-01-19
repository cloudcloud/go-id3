package main

import (
	"context"
	"errors"
	"os"

	"github.com/cloudcloud/go-id3"
	"github.com/urfave/cli/v3"
)

var readCmdDescription = `
Read comprehensive information about a specific file.

For example:
	go-id3 read /var/filename.mp3
`

var readCmd = &cli.Command{
	Name:        "read",
	Usage:       "Display information from a specific file",
	Description: readCmdDescription,
	Arguments: []cli.Argument{
		&cli.StringArg{
			Name: "filename",
		},
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}},
		&cli.StringFlag{Name: "format", Aliases: []string{"f"}},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		input := cmd.StringArg("filename")
		format := cmd.String("format")

		if input == "" {
			return errors.New("a filename is required for `read`")
		}

		debug := cmd.Bool("debug")
		output, err := id3.Process(ctx, input, debug)
		if err != nil {
			return err
		}

		output.PrettyPrint(os.Stdout, format)

		return nil
	},
}
