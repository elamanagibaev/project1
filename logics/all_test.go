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
		{"ProcessUp (with punctuation)", ProcessUp, "elaman, (up)", "ELAMAN,"},
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
		input string
		base  string
		punct string
	}{
		{"elaman,", "elaman", ","},
		{"word.", "word", "."},
		{"hello!", "hello", "!"},
		{"simple", "simple", ""},
		{"", "", ""},
	}

	for _, test := range tests {
		base, punct := CleanWord(test.input)
		if base != test.base || punct != test.punct {
			t.Errorf(" CleanWord(\"%s\"): expected base=\"%s\", punct=\"%s\" → got base=\"%s\", punct=\"%s\"",
				test.input, test.base, test.punct, base, punct)
		}
	}
}
