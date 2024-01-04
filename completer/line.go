package completer

type Line struct {
	data []rune
}

func makeLine(line []rune) Line {
	return Line{line}
}

func (line *Line) words() [][]rune {
	var words [][]rune = make([][]rune, 0)
	var word []rune = make([]rune, 0)
	for _, c := range line.data {
		if c == ' ' {
			words = append(words, word)
			word = make([]rune, 0)
		} else {
			word = append(word, c)
		}
	}
	words = append(words, word)
	return words
}

func (line *Line) previousWord(pos int) []rune {
	// the result
	var word []rune = make([]rune, 0)

	// get the last space position
	var lastSpaceFound bool = false
	for i := pos - 1; i >= 0; i-- {
		if line.data[i] == ' ' {
			pos = i
			lastSpaceFound = true
			break
		}
	}

	if !lastSpaceFound {
		return []rune("")
	}

	for i := pos - 1; i >= 0; i-- {
		if line.data[i] == ' ' {
			break
		}
		word = append(word, line.data[i])
	}
	// reverse the word
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return word
}

func (line *Line) currentWord(pos int) []rune {
	// the result
	var word []rune = make([]rune, 0)

	for i := pos - 1; i >= 0; i-- {
		if line.data[i] == ' ' {
			break
		}
		word = append(word, line.data[i])
	}
	// reverse the word
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return word
}
