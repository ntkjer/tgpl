package isanagram

import (
	"strings"
)

func isAnagram(stringA, stringB string) bool {
	shortString, longString := stringA, stringB
	if len(stringA) > len(stringB) {
		shortString = stringB
		longString = stringA
	}

	n := len(longString)
	var result bool
	for i := 0; i < n; n++ {
		if strings.Contains(longString, shortString[i:]) {
			result = true
		} else {
			result = false
		}
	}
	return result
}
