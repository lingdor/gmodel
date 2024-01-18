package stringutil

import "unicode"

func Contains(strs []string, str string) int {
	for i, item := range strs {
		if item == str {
			return i
		}
	}
	return -1
}

func UpperFrist(str string) string {

	if str == "" {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
