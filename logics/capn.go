package logics

import (
	"strconv"
	"strings"
)

func ProcessCapN(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap,") && strings.HasSuffix(words[i], ")") {
			nStr := strings.TrimSuffix(strings.TrimPrefix(words[i], "(cap,"), ")")
			n, err := strconv.Atoi(strings.TrimSpace(nStr))
			if err == nil && i-n >= 0 {
				for j := i - n; j < i; j++ {
					word := words[j]
					if len(word) > 0 {
						words[j] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
					}
				}
				words[i] = ""
			}
		}
	}

	return strings.Join(words, " ")
}
