package completer

import (
	"slices"
	"strings"
)

type OptionCompleter struct {
	OptionName             []rune
	OptionValues           [][]rune
	OptionValueCompleterFn func() [][]rune
	MultipleValues         bool
}

type CmdCompleter struct {
	CmdName          []rune
	FlagNames        [][]rune
	OptionCompleters []OptionCompleter
}

type CmdCompleters struct {
	Completers []CmdCompleter
}

func (cmdCompleters *CmdCompleters) AddCmdCompleter(cmdCompleter CmdCompleter) {
	cmdCompleters.Completers = append(cmdCompleters.Completers, cmdCompleter)
}

func (cmdCompleters *CmdCompleters) Do(line []rune, pos int) ([][]rune, int) {
	l := makeLine(line)
	words := l.words()

	var candidates [][]rune = make([][]rune, 0)

	if len(words) == 1 {
		cmdName := words[0]
		for _, cmdCompleter := range cmdCompleters.Completers {
			if strings.HasPrefix(string(cmdCompleter.CmdName), string(cmdName)) {
				candidates = append(candidates, cmdCompleter.CmdName[len(cmdName):])
			}
		}
		if len(candidates) > 0 {
			return candidates, len(cmdName)
		}
	} else {
		cmdName := words[0]

		for _, cmdCompleter := range cmdCompleters.Completers {
			if string(cmdCompleter.CmdName) == string(cmdName) {
				return cmdCompleter.Do(line, pos)
			}
		}
	}

	return nil, 0
}

func (cmdCompleter *CmdCompleter) Do(line []rune, pos int) ([][]rune, int) {
	l := makeLine(line)
	words := l.words()
	currentWord := l.currentWord(pos)
	previousWord := l.previousWord(pos)

	var candidates [][]rune = make([][]rune, 0)

	for _, optionCompleter := range cmdCompleter.OptionCompleters {
		optionName := optionCompleter.OptionName

		// if the previousWord is the optionName,
		// then the candidates are made from the optionValueCompleterFn
		if string(optionName) == string(previousWord) {
			for _, optionValue := range optionCompleter.OptionValues {
				// if currentWord does not have a prefix of optionValue, then skip it
				if !strings.HasPrefix(string(optionValue), string(currentWord)) {
					continue
				}
				candidates = append(candidates, optionValue[len(currentWord):])
			}

			if optionCompleter.OptionValueCompleterFn != nil {
				for _, optionValue := range optionCompleter.OptionValueCompleterFn() {
					// if currentWord does not have a prefix of optionValue, then skip it
					if !strings.HasPrefix(string(optionValue), string(currentWord)) {
						continue
					}
					candidates = append(candidates, optionValue[len(currentWord):])
				}
			}
			return candidates, len(currentWord)
		}
	}

	for _, optionCompleter := range cmdCompleter.OptionCompleters {
		optionName := optionCompleter.OptionName
		multipleValues := optionCompleter.MultipleValues

		// if the optionName is in the words, then skip it
		if !multipleValues && runesSliceContains(words, optionName) && string(optionName) != string(currentWord) {
			continue
		}

		if strings.HasPrefix(string(optionName), string(currentWord)) {
			// add the optionName to the candidates
			candidates = append(candidates, optionName[len(currentWord):])
		}
	}

	for _, flagName := range cmdCompleter.FlagNames {
		// if the optionName is in the words, then skip it
		if runesSliceContains(words, flagName) && string(flagName) != string(currentWord) {
			continue
		}

		if strings.HasPrefix(string(flagName), string(currentWord)) {
			// add the flagName to the candidates
			candidates = append(candidates, flagName[len(currentWord):])
		}
	}

	return candidates, len(currentWord)

}

func runesSliceContains(words [][]rune, optionName []rune) bool {
	for _, word := range words {
		if slices.Equal(optionName, word) {
			return true
		}
	}
	return false
}
