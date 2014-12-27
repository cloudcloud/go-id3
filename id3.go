package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Command struct {
	Run                    func(args []string)
	UsageLine, Short, Long string
}

type LoggedError struct {
	error
}

var commands = []*Command{
	readCmd,
}

const usageTemplate = `usage: go-id3 command [arguments]
The commands are:
{{range .}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "go-id3 help [command]" for more information.
`

const helpTemplate = `usage: go-id3 {{.UsageLine}}
{{.Long}}
`

func main() {
	flag.Usage = func() { usage(1) }
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || args[0] == "help" {
		if len(args) == 1 {
			usage(0)
		}
		if len(args) > 1 {
			for _, cmd := range commands {
				if cmd.Name() == args[1] {
					tmpl(os.Stdout, helpTemplate, cmd)
					return
				}
			}
		}
		usage(2)
	}

	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(LoggedError); !ok {
				panic(err)
			}
			os.Exit(1)
		}
	}()

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Run(args[1:])
			return
		}
	}

	errorf("unknown command [%q]\nRun 'go-id3 help' for usage.\n", args[0])
}

func errorf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stderr, format, args...)
	panic(LoggedError{})
}

func usage(exit int) {
	tmpl(os.Stderr, usageTemplate, commands)
	os.Exit(exit)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func (cmd *Command) Name() string {
	name := cmd.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}
