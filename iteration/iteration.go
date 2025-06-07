package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var stringBuilder strings.Builder

	for range repeatCount {
		stringBuilder.WriteString(character)
	}

	return stringBuilder.String()
}
