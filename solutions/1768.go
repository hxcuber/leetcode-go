package solutions

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	wordLess := word1
	wordMore := word2
	if len(word1) > len(word2) {
		wordLess = word2
		wordMore = word1
	}

	sb := strings.Builder{}
	for i := 0; i < len(wordLess); i++ {
		sb.WriteRune(rune(word1[i]))
		sb.WriteRune(rune(word2[i]))
	}
	sb.WriteString(wordMore[len(wordLess):])
	return sb.String()
}
