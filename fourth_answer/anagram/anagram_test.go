package anagram_test

import (
	"github.com/muhammadaser/stockbit/fourth_answer/anagram"
	"reflect"
	"testing"
)

func TestGroupOfAnagramString(t *testing.T) {
	// crate test case
	testCase := map[string]struct {
		Input    []string
		Expected [][]string
	}{
		"Case Empty": {
			Input:    []string{},
			Expected: nil,
		},
		"Case only one data": {
			Input: []string{"kata"},
			Expected: [][]string{
				{"kata"},
			},
		},
		"Case multiple data": {
			Input: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
			Expected: [][]string{
				{"kita", "atik", "tika"},
				{"aku", "kua"},
				{"kia"},
				{"makan"},
			},
		},
	}

	for caseName, test := range testCase {
		// run function to test
		actual := anagram.GroupOfAnagramString(test.Input)

		// for empty case
		if test.Expected == nil {
			if actual != nil {
				t.Errorf("%s : expected %+v , actual %+v", caseName, test.Expected, actual)
			}
			// continue to next case
			continue
		}

		// check test case
		if !reflect.DeepEqual(actual, test.Expected) {
			t.Errorf("%s : expected %+v , actual %+v", caseName, test.Expected, actual)
		}
	}
}
