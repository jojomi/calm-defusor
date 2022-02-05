package communication

import "fmt"

func SafeWord(input string, wordList map[string]string) string {
	if explanation, ok := wordList[input]; ok {
		return fmt.Sprintf("%s (%s)", input, explanation)
	}
	return input
}
