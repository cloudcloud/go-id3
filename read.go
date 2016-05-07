package main

import (
	"fmt"

	"github.com/cloudcloud/go-id3/file"
)

var readCmd = &Command{
	UsageLine: "read [options] [filename]",
	Short:     "Display information from a specific file",
	Long: `
Read comprehensive information about a specific file.

Options:


For example:
	go-id3 read /var/filename.mp3
`,
}

func init() {
	readCmd.Run = readRun
}

func readRun(args []string) {
	if len(args) < 1 {
		errorf("No filename provided")
	}

	defer func() {
		if r := recover(); r != nil {
			// need to do some better handling of big problems
			fmt.Println("Encountered a panic, please resolve:\n\t", r)
		}
	}()

	f := &file.File{Filename: args[0], Debug: isDebug}
	defer f.CleanUp()

	fmt.Println(f.Process().PrettyPrint())
}
