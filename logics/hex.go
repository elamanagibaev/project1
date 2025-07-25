package logics

import (
	"strconv"
	"strings"
)

func ProcessHex(line string) string {
	if !strings.Contains(line, "(hex)") {
		return line
	}
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexStr := words[i-1]

			n, err := strconv.ParseInt(hexStr, 0, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, "")
}
