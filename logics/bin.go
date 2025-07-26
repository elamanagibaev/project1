package logics

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessBin(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			num, err := strconv.Atoi(words[i-1])
			if err == nil {
				words[i-1] = fmt.Sprintf("0b%b", num)
				words[i] = ""
			}
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
