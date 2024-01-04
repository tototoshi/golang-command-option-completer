package completion

import (
	"time"

	"example.com/demo/completer"
)

var Curl completer.CmdCompleter = completer.CmdCompleter{
	CmdName:   []rune("curl"),
	FlagNames: [][]rune{[]rune("--version"), []rune("--help")},
	OptionCompleters: []completer.OptionCompleter{
		{
			OptionName: []rune("--url"),
			OptionValues: [][]rune{
				[]rune("https://httpbin.org/get"),
				[]rune("https://httpbin.org/post"),
			},
		},
		{
			OptionName: []rune("--request"),
			OptionValues: [][]rune{
				[]rune("GET"),
				[]rune("POST"),
			},
		},
		{
			OptionName:     []rune("--data-urlencode"),
			MultipleValues: true,
		},
		{
			OptionName:     []rune("--url-query"),
			MultipleValues: true,
		},
		{
			OptionName: []rune("-o"),
			OptionValueCompleterFn: func() [][]rune {
				timestamp := time.Now().Format("20060102150405")
				filename := "output_" + timestamp + ".txt"
				return [][]rune{[]rune(filename)}
			},
		},
	},
}
