package main

import (
	"encoding/json"
	"fmt"
	"github.com/stockbit/fourth_answer/anagram"
	"io/ioutil"
	"log"
)

func main() {
	// read list of string from file
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("list string : %s \n", string(content))

	// set variable for list of string
	var listOfStr []string

	// unmarshal list of string
	err = json.Unmarshal(content, &listOfStr)
	if err != nil {
		log.Fatal(err)
	}

	// do grouping for anagram string
	result := anagram.GroupOfAnagramString(listOfStr)

	// print result of anagram string
	fmt.Println(result)
}
