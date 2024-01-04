package cmd

import (
	"os"
	"os/exec"

	"example.com/demo/cmd/completion"
	"example.com/demo/completer"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo",
	Long:  "demo",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "> ",
		AutoComplete: &completer.CmdCompleters{
			Completers: []completer.CmdCompleter{
				completion.Curl,
				completion.Ls,
			},
		},
		HistoryFile: "/tmp/demo_history.txt",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		// exec "line" as shell command and print output
		cmd := exec.Command("/bin/sh", "-c", line)
		cmd.Env = os.Environ()

		out, err := cmd.CombinedOutput()

		rl.SaveHistory(line)

		if err != nil {
			println("Error: " + err.Error())
			continue
		}
		print(string(out))
	}
}

func Execute() {
	rootCmd.Execute()
}
