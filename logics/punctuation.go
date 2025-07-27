package logics

import "strings"

func CleanWord(word string) (string, string) {
	if len(word) == 0 {
		return "", ""
	}
	last := word[len(word)-1]
	if strings.ContainsRune(",.!?:;'", rune(last)) {
		return word[:len(word)-1], string(last)
	}
	return word, ""
}
