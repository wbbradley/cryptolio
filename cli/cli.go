package cli

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("cryptolio", "A crypto financial planning tool")
	commands = map[string]func(){}
)

func Register(name, desc string, dispatch func()) *kingpin.CmdClause {
	cmd := app.Command(name, desc)
	commands[cmd.FullCommand()] = dispatch
	return cmd
}

func Dispatch() {
	if dispatch, ok := commands[kingpin.MustParse(app.Parse(os.Args[1:]))]; ok {
		dispatch()
	}
}
