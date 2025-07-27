package logics

import "strings"

func CleanWord(word string) (string, string) {
	if word == "" {
		return "", ""
	}

	punctuation := ",.!?:;\"'’…—–•™©!?”“"
	runes := []rune(word)
	end := len(runes)

	for end > 0 && strings.ContainsRune(punctuation, runes[end-1]) {
		end--
	}

	base := string(runes[:end])
	punct := string(runes[end:])

	return base, punct
}
