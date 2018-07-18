package stringops

import (
	"bytes"
	"strconv"
)

// Compress shortens a string: AAABBC = A3B2C1
func Compress(s string) string {
	if s == "" {
		return ""
	}
	var prev string
	var buffer bytes.Buffer
	var count int
	var v rune
	for _, v = range s {
		if prev == "" {
			prev = string(v)
			continue
		} else if prev != string(v) && prev != "" {
			prev, count = bufWrite(&buffer, prev, count, v)
		} else if prev == string(v) {
			count += 1
		}
	}
	bufWrite(&buffer, prev, count, v)

	return buffer.String()

}

func bufWrite(b *bytes.Buffer, prev string, count int, v rune) (string, int) {
	b.WriteString(prev)
	b.WriteString(strconv.Itoa(count))
	prev = string(v)
	count = 1
	return prev, count
}
