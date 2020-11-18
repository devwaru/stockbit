package utils_test

import (
	"github.com/muhammadaser/stockbit/third_answer/utils"
	"testing"
)

func TestFindFirstStringInBracket(t *testing.T) {
	// crate test case
	testCase := map[string]struct {
		Input    string
		Expected string
	}{
		"Case Empty string": {
			Input:    "",
			Expected: "",
		},
		"Case no bracket": {
			Input:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
			Expected: "",
		},
		"Case multiple bracket": {
			Input:    "Lorem (Ipsum) is simply dummy (text of the printing) and typesetting industry.",
			Expected: "Ipsum",
		},
		"Case string in bracket have white space": {
			Input:    "Lorem (Ipsum is )simply dummy (text of the printing) and typesetting industry.",
			Expected: "Ipsum is",
		},
	}

	for caseName, test := range testCase {
		// run function to test
		actual := utils.FindFirstStringInBracket(test.Input)

		// check test case
		if actual != test.Expected {
			t.Errorf("%s : expected %s , actual %s", caseName, test.Expected, actual)
		}
	}
}
