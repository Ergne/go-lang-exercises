package kata

import (
	"strings"
)

func DuplicateEncode(word string) string {
	var convertedWord []string
	word = strings.ToLower(word)
	for i, char := range word {
		for icompare, comparechar := range word {
			//fmt.Println(i, icompare, char, comparechar)
			if char == comparechar && i != icompare {
				convertedWord = append(convertedWord, ")")
				break
			}
			if icompare == len(word)-1 {
				convertedWord = append(convertedWord, "(")
				break
			}
		}
	}
	return strings.Join(convertedWord, "")
}
