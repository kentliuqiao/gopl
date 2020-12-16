package e7_3_test

import (
	"gopl/e7_3"
	"testing"
)

func TestSort(t *testing.T) {

	tree := e7_3.Add(nil, 2)
	tree = e7_3.Add(tree, 3)
	tree = e7_3.Add(tree, 1)

	s := tree.String()
	expected := "[2, 1, 3, ]"
	if s != expected {
		t.Errorf("String() returned %s. expected: %s", s, expected)
	}
}
