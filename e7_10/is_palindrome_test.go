package e7_10

import (
	"sort"
	"testing"
)

func TestStringSlice(t *testing.T) {
	if !IsPalindrome(sort.StringSlice([]string{"aaa"})) {
		t.Errorf("['aaa'] is palindrome")
	}

	if !IsPalindrome(sort.StringSlice([]string{"aaa", "bbb", "aaa"})) {
		t.Errorf("['aaa', 'bbb', 'aaa'] is palindrome")
	}

	if !IsPalindrome(sort.StringSlice([]string{"aaa", "aaa"})) {
		t.Errorf("['aaa', 'aaa'] is palindrome")
	}

	if !IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "bbb", "aaa"})) {
		t.Errorf("['aaa', 'bbb' 'bbb', 'aaa'] is palindrome")
	}

	if IsPalindrome(sort.StringSlice([]string{"aaa", "bbb"})) {
		t.Errorf("['aaa', 'bbb'] is not palindrome")
	}

	if IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "ccc", "ddd"})) {
		t.Errorf("['aaa', 'bbb' 'ccc', 'ddd'] is not palindrome")
	}

	if IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "ccc", "aaa"})) {
		t.Errorf("['aaa', 'bbb' 'ccc', 'aaa'] is not palindrome")
	}

	if IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "bbb", "bbb"})) {
		t.Errorf("['aaa', 'bbb' 'bbb', 'bbb'] is not palindrome")
	}
}

func TestIntSlice(t *testing.T) {
	if !IsPalindrome(sort.IntSlice([]int{1})) {
		t.Errorf("[1] is palindrome")
	}

	if !IsPalindrome(sort.IntSlice([]int{1, 1})) {
		t.Errorf("[1, 1] is palindrome")
	}

	if !IsPalindrome(sort.IntSlice([]int{1, 1, 1})) {
		t.Errorf("[1, 1, 1] is palindrome")
	}

	if !IsPalindrome(sort.IntSlice([]int{1, 2, 1})) {
		t.Errorf("[1, 2, 1] is palindrome")
	}

	if !IsPalindrome(sort.IntSlice([]int{1, 2, 2, 1})) {
		t.Errorf("[1, 2, 2, 1] is palindrome")
	}

	if IsPalindrome(sort.IntSlice([]int{1, 2})) {
		t.Errorf("[1, 2] is not palindrome")
	}

	if IsPalindrome(sort.IntSlice([]int{1, 2, 3, 4})) {
		t.Errorf("[1, 2, 3, 4] is not palindrome")
	}

	if IsPalindrome(sort.IntSlice([]int{1, 2, 3, 1})) {
		t.Errorf("[1, 2, 3, 1] is not palindrome")
	}

	if IsPalindrome(sort.IntSlice([]int{1, 2, 2, 2})) {
		t.Errorf("[1, 2, 2, 2] is not palindrome")
	}
}
