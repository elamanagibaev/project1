package logics

import (
	"strings"
)

func ProcessLow(line string) string {
	if !strings.Contains(line, "(low)") {
		return line
	}

	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			base, punct := CleanWord(words[i-1])
			words[i-1] = strings.ToLower(base) + punct
			words[i] = ""
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
