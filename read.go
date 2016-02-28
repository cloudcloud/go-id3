package main

import "github.com/cloudcloud/go-id3/id3"

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
	if len(args) < 1 {
		errorf("No filename provided")
	}

	filename = args[0]
	i, err := id3.New(filename)
	if err != nil {
		errorf("File does not exist")
	}

	if isDebug {
		i.SetDebug()
	}

	i.Process()
	i.PrettyPrint()
}
