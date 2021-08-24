package translator

import (
	"strings"
)

const lineThreshold = 77

// split split long text to small parts by blank or newline(\n)
func split(long string, threshold int) []string {
	return split0(long, threshold, 0)
}

// split split long text to small parts by blank or newline(\n)
func split0(long string, threshold int, depth int) []string {
	length := len(long)
	if index := strings.Index(long, "\n"); index >= 0 && index < length-1 {
		left := long[:index+1]
		right := long[index+1:]
		prefix := []string{left}
		if depth == 0 {
			prefix = []string{"", left}
		}
		return append(prefix, split0(right, threshold, depth+1)...)
	}
	if length <= threshold {
		return []string{long}
	}
	if index := strings.LastIndex(long, " "); index >= 0 && index < length-1 {
		left := long[:index+1]
		right := long[index+1:]
		return append(split0(left, threshold, depth+1), right)
	}
	return []string{long}
}
