package utils

import (
	"regexp"
	"strings"
)

func capitalizeSentences(text string) string {
	// Split the text into individual sentences using regex
	re := regexp.MustCompile(`[.?!]\s+`)
	sentences := re.Split(text, -1)

	// Capitalize the first character of each sentence
	capitalizedSentences := make([]string, len(sentences))
	for i, sentence := range sentences {
		capitalizedSentences[i] = strings.ToUpper(string(sentence[0])) + strings.ToLower(sentence[1:])
	}
	capitalizedText := strings.Join(capitalizedSentences, " ")

	return capitalizedText
}

func RemoveDuplicates(strings []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, str := range strings {
		if !encountered[str] {
			encountered[str] = true
			result = append(result, str)
		}
	}

	return result
}
