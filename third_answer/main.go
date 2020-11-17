package main

import (
	"fmt"
	"regexp"
)

func main() {
	// get string from user via terminal
	fmt.Println("Enter the string : ")

	// set variable 'str' to store input from user
	var str string

	// scan user input
	fmt.Scanln(&str)

	// find first string in bracket from user input
	firstString := findFirstStringInBracket(str)

	// show result
	if len(firstString) > 0 {
		fmt.Println("First string in Bracket is " + firstString)
	} else {
		fmt.Println("String in Bracket not found")
	}
}

func findFirstStringInBracket(str string) string {
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

	// return and trim bracket from string
	return firstStringWithBracket[1:][:len(firstStringWithBracket)-2]
}
