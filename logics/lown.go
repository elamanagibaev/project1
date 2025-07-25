package logics

import (
	"strconv"
	"strings"
)

func ProcessLowN(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(low,") && strings.HasSuffix(words[i], ")") {
			nStr := strings.TrimSuffix(strings.TrimPrefix(words[i], "(low,"), ")")
			n, err := strconv.Atoi(strings.TrimSpace(nStr))
			if err == nil && i-n >= 0 {
				for j := i - n; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
				words[i] = ""
			}
		}
	}

	return strings.Join(words, " ")
}
