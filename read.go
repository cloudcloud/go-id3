package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cloudcloud/go-id3/file"
)

var readCmd = &Command{
	UsageLine: "read [filename]",
	Short:     "Display information from a specific file",
	Long: `
Read comprehensive information about a specific file.

For example:
	go-id3 read /var/filename.mp3
`,
}

func init() {
	readCmd.Run = readProcess
}

func readProcess(args []string, o io.Writer) {
	if len(args) < 1 {
		fmt.Fprintf(o, "No filename provided")
		return
	}

	defer catcher(os.Stderr)

	f := &file.File{Filename: args[0], Debug: isDebug}
	handle, err := os.Open(f.Filename)
	if err != nil {
		fmt.Fprintf(o, "Unable to open the file [%s]", f.Filename)
		return
	}
	defer handle.Close()

	f.Process(handle).PrettyPrint(o, outFormat)
}
