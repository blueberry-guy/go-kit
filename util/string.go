package util

import "strings"

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
