package logics

import (
	"strings"
)

func ProcessCap(line string) string {
	if !strings.Contains(line, "(cap)") {
		return line
	}

	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			base, punct := CleanWord(words[i-1])
			words[i-1] = strings.Title(base) + punct
			words[i] = ""
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
