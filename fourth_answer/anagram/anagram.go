package anagram

import "sort"

func GroupOfAnagramString(sliceOfStr []string) [][]string {
	var result [][]string

	// return empty result if sliceOfStr is empty
	if len(sliceOfStr) <= 0 {
		return result
	}

	// create hash table to group string by sum of ascii number.
	tempHashTable := make(map[int][]string)
	for _, str := range sliceOfStr {
		sumAscii := sumOfAsciiString(str)
		tempSliceOfStr, ok := tempHashTable[sumAscii]
		if ok {
			tempHashTable[sumAscii] = append(tempSliceOfStr, str)
		} else {
			tempHashTable[sumAscii] = []string{str}
		}
	}

	// append into result variable
	result = make([][]string, 0)
	for _, val := range tempHashTable {
		result = append(result, val)
	}

	// sort base on len of anagram group
	sort.Slice(result, func(i, j int) bool { return len(result[i]) > len(result[j]) })

	// return result
	return result
}

func sumOfAsciiString(str string) int {
	var sum int

	// return 0 if len of str is zero
	if len(str) <= 0 {
		return 0
	}

	// do sum for ascii number.
	for _, runeOfStr := range []rune(str) {
		sum += int(runeOfStr)
	}

	return sum
}
