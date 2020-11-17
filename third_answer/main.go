package main

import (
	"fmt"
	"github.com/stockbit/third_answer/utils"
)

func main() {
	// get string from user via terminal
	fmt.Println("Enter the string : ")

	// set variable 'str' to store input from user
	var str string

	// scan user input
	fmt.Scanln(&str)

	// find first string in bracket from user input
	firstString := utils.FindFirstStringInBracket(str)

	// show result
	if len(firstString) > 0 {
		fmt.Println("First string in Bracket is " + firstString)
	} else {
		fmt.Println("String in Bracket not found")
	}
}
