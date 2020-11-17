package utils

import (
	"regexp"
	"strings"
)

func FindFirstStringInBracket(str string) string {
	// return empty string if str len is zero or less
	if len(str) <= 0 {
		return ""
	}

	// find first string inside bracket with regex
	re := regexp.MustCompile(`\((.*?)\)`)
	firstStringWithBracket := re.FindString(str)

	// return empty string if string with bracket not exist
	if len(firstStringWithBracket) <= 0 {
		return ""
	}

	// trim bracket from string
	resultString := firstStringWithBracket[1:][:len(firstStringWithBracket)-2]

	// return and trim white space
	return strings.Trim(resultString, " ")
}
