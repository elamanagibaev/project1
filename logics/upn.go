package logics

import (
	"strconv"
	"strings"
)

func ProcessUpN(line string) string {
	words := strings.Fields(line)

	for i := range words {
		if strings.HasPrefix(words[i], "(up,") && strings.HasSuffix(words[i], ")") {
			marker := strings.TrimPrefix(words[i], "(up,")
			marker = strings.TrimSuffix(marker, ")")
			nStr := strings.TrimSpace(marker)

			n, err := strconv.Atoi(nStr)
			if err != nil {
				continue
			}

			start := i - n
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				base, punct := CleanWord(words[j])
				words[j] = strings.ToUpper(base) + punct
			}

			words[i] = ""
		}
	}

	var filtered []string
	for _, w := range words {
		if w != "" {
			filtered = append(filtered, w)
		}
	}

	return strings.Join(filtered, " ")
}
