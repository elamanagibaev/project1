package logics

import (
	"strconv"
	"strings"
)

func ProcessBin(line string) string {
	if !strings.Contains(line, "(bin)") {
		return line
	}
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binStr := words[i-1]

			n, err := strconv.ParseInt(binStr, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, " ")
}
