package orm

import (
	"strings"
	"unicode"
)

func ToDbName(fieldName string, gmodelTag string) string {

	tagVals := strings.Split(gmodelTag, ",")
	if len(tagVals) > 0 && strings.TrimSpace(tagVals[0]) != "" {
		return tagVals[0]
	}
	newName := make([]rune, 0, len(fieldName))
	for i, r := range []rune(fieldName) {
		if i > 0 {
			if unicode.IsUpper(r) {
				newName = append(newName, '_', unicode.ToLower(r))
				continue
			}
		} else if i == 0 {
			newName = append(newName, unicode.ToLower(r))
			continue
		}
		newName = append(newName, r)
	}
	return string(newName)
}
