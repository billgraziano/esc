package embed

import "strings"

func replaceNewLines(src string) string {
	return strings.Replace(src, "\r\n", "\n", -1)
}
