package completion

import "example.com/demo/completer"

var Ls completer.CmdCompleter = completer.CmdCompleter{
	CmdName:   []rune("ls"),
	FlagNames: [][]rune{[]rune("-l"), []rune("-a"), []rune("-h")},
}
