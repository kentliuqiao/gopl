package e7_10

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, l := 0, s.Len(); i < l/2; i++ {
		if s.Less(i, l-i-1) || s.Less(l-i-1, i) {
			return false
		}
	}

	return true
}
