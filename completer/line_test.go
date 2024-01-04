package completer

import "testing"

func Test_words(t *testing.T) {
	a := makeLine([]rune("abc def"))
	result := a.words()
	if len(result) != 2 {
		t.Error("Expected 2, got ", len(result))
	}
	if string(result[0]) != "abc" {
		t.Error("Expected abc, got ", string(result[0]))
	}
	if string(result[1]) != "def" {
		t.Error("Expected def, got ", string(result[1]))
	}
}

func Test_previous_word(t *testing.T) {
	a := makeLine([]rune("abc def"))
	result := a.previousWord(7)
	if string(result) != "abc" {
		t.Error("Expected abc, got ", string(result))
	}

	result = a.previousWord(3)
	if string(result) != "" {
		t.Error("Expected empty string, got ", string(result))
	}
}

func Test_current_word(t *testing.T) {
	a := makeLine([]rune("abc def"))
	result := a.currentWord(7)
	if string(result) != "def" {
		t.Error("Expected def, got ", string(result))
	}

	result = a.currentWord(3)
	if string(result) != "abc" {
		t.Error("Expected abc, got ", string(result))
	}
}
