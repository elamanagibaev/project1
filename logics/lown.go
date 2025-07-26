package logics

import (
	"strconv"
	"strings"
)

func ProcessLowN(line string) string {
	words := strings.Fields(line)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") && strings.Contains(words[i], "low") {
			marker := strings.TrimPrefix(words[i], "(")
			marker = strings.TrimSuffix(marker, ")")
			parts := strings.Split(marker, ",")
			if len(parts) != 2 {
				continue
			}
			nStr := strings.TrimSpace(parts[1])
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
				words[j] = strings.ToLower(base) + punct
			}
			words[i] = ""
		}
	}
	return strings.TrimSpace(strings.Join(words, " "))
}
