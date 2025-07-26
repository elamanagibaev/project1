package logics

import (
	"strings"
)

func ProcessUp(line string) string {
	if !strings.Contains(line, "(up)") {
		return line
	}

	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			base, punct := CleanWord(words[i-1])
			words[i-1] = strings.ToUpper(base) + punct
			words[i] = ""
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
