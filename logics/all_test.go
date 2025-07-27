package logics

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) string
		input    string
		expected string
	}{
		{"ProcessUp", ProcessUp, "elaman (up)", "ELAMAN"},
		{"ProcessUp", ProcessUp, "elaman, (up)", "ELAMAN,"},
		{"ProcessLow", ProcessLow, "GOLANG (low)", "golang"},
		{"ProcessCap", ProcessCap, "hello (cap)", "Hello"},
		{"ProcessUpN", ProcessUpN, "one two (up,2)", "ONE TWO"},
		{"ProcessLowN", ProcessLowN, "HELLO WORLD (low,2)", "hello world"},
		{"ProcessCapN", ProcessCapN, "again here (cap,2)", "Again Here"},
		{"ProcessHex", ProcessHex, "1337 (hex)", "0x539"},
		{"ProcessBin", ProcessBin, "42 (bin)", "0b101010"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input)
			if result != tt.expected {
				t.Errorf(" %s: expected \"%s\", got \"%s\"", tt.name, tt.expected, result)
			}
		})
	}
}

func TestPunctuation(t *testing.T) {
	tests := []struct {
		input         string
		expectedBase  string
		expectedPunct string
	}{
		{"hello,", "hello", ","},
		{"world.", "world", "."},
		{"wow!", "wow", "!"},
		{"maybe?", "maybe", "?"},
		{"okay:", "okay", ":"},
		{"yes;", "yes", ";"},
		{"stop\"", "stop", "\""},
		{"cool'", "cool", "'"},
		{"he’s’", "he’s", "’"},
		{"fine", "fine", ""},
		{"", "", ""},

		{"it's", "it's", ""},
		{"it's'", "it's", "'"},
		{"I'm", "I'm", ""},
		{"I'm'", "I'm", "'"},
		{"they’ll", "they’ll", ""},
		{"won’t'", "won’t", "'"},
		{"Elaman’", "Elaman", "’"},
		{"don't.", "don't", "."},

		{"'hello'", "'hello", "'"},
		{"\"hello\"", "\"hello", "\""},
		{"'I’m'", "'I’m", "'"},
		{"“great.”", "“great", ".”"},

		{"hello...", "hello", "..."},
		{"wow!?", "wow", "!?"},
		{"wait?!", "wait", "?!"},
		{"sure!!!", "sure", "!!!"},
		{"no!!!", "no", "!!!"},
		{"hello??", "hello", "??"},
		{"hmm...", "hmm", "..."},
		{"amazing?!", "amazing", "?!"},

		{"go,'", "go", ",'"},
		{"she's!'", "she's", "!'"},
		{"we're?!", "we're", "?!"},
		{"I’M!", "I’M", "!"},
		{"yes!'", "yes", "!'"},

		{"pause…", "pause", "…"},
		{"dash—", "dash", "—"},
		{"dash–", "dash", "–"},
		{"dot•", "dot", "•"},
		{"brand™", "brand", "™"},
		{"copyright©", "copyright", "©"},

		{"'", "", "'"},
		{"’", "", "’"},
		{"\"", "", "\""},
		{".", "", "."},
		{"!", "", "!"},
		{";", "", ";"},

		{"rock'n'roll", "rock'n'roll", ""},
		{"it's'", "it's", "'"},
		{"they’re", "they’re", ""},
		{"O’Neil!", "O’Neil", "!"},

		{"This is great!", "This is great", "!"},
		{"No way?!", "No way", "?!"},
		{"I said: 'Go!'", "I said: 'Go", "!'"},
		{"That's John's car.", "That's John's car", "."},
		{"He replied, “Okay.”", "He replied, “Okay", ".”"},

		{"'Yes!'", "'Yes", "!'"},
		{"‘Right?’", "‘Right", "?’"},
		{"'No...'", "'No", "...'"},
		{"\"Absolutely!\"", "\"Absolutely", "!\""},
		{"‘Indeed.’", "‘Indeed", ".’"},
	}

	for _, tt := range tests {
		base, punct := CleanWord(tt.input)
		if base != tt.expectedBase || punct != tt.expectedPunct {
			t.Errorf("CleanWord(%q): expected base=%q, punct=%q → got base=%q, punct=%q",
				tt.input, tt.expectedBase, tt.expectedPunct, base, punct)
		}
	}
}
