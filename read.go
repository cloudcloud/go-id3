package main

import (
	"github.com/cloudcloud/go-id3/src/id3"
)

var readCmd = &Command{
	UsageLine: "read [filename]",
	Short:     "Display information from a specific file",
	Long: `
Read comprehensive information about a specific file.

Options:
	

For example:
	go-id3 read /var/filename.mp3
`,
}

var (
	filename string
)

func init() {
	readCmd.Run = readRun
}

func readRun(args []string) {
	if len(args) != 1 {
		errorf("No filename provided")
	}

	filename = args[0]
	i := id3.New(filename)
	i.Process()
	i.PrettyPrint()
}
