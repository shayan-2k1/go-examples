package findian

import (
	"strings"
)

func Find(s string) bool {
	if strings.HasPrefix(s, "i") {
		if strings.HasSuffix(s, "n") {
			if strings.Contains(s, "a") {
				return true
			}
		}
	}
	return false
}
