package normalize

import (
	"bytes"
	"regexp"
)

func DoString(phoneNumber string) string {
	var buf bytes.Buffer

	for _, ch := range phoneNumber {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func DoRegex(p string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(p, "")
}
