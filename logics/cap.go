package logics

import "strings"

func ProcessCap(line string) string {
	if !strings.Contains(line, "(cap)") {
		return line
	}
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			word := words[i-1]
			if len(word) > 0 {
				words[i-1] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
