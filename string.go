package utils

import (
	"strings"
)

// SplitAfter is plus of strings.SplitAfter.
func SplitAfter(s string, sep ...string) []string {
	c := 0
	for _, sp := range sep {
		c += strings.Count(s, sp)
	}
	ret := []string{s}
	if c == 0 {
		return ret
	}
	for _, sp := range sep {
		n := []string{}
		for _, r := range ret {
			for _, a := range strings.SplitAfter(r, sp) {
				aa := strings.TrimSpace(a)
				if aa != "" {
					n = append(n, aa)
				}
			}
		}
		ret = n
	}
	return ret
}
