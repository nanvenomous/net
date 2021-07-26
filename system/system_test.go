package system

import (
	"testing"
)

func slicesEqual(t *testing.T, one []string, two []string) {
	if (one == nil) || (two == nil) {
		t.Error("an array was nil")
	}
	if len(one) != len(two) {
		t.Error("array lengths not equal")
	}
	for i := range one {
		if one[i] != two[i] {
			t.Errorf("inequality at %d", i)
		}
	}
}

func TestPrependArgument(t *testing.T) {
	firstString := "first_string"
	secondString := "second_string"
	arr := []string{secondString}
	expectedArr := []string{firstString, secondString}
	arr = PrependArgument(firstString, arr)
	slicesEqual(t, arr, expectedArr)
}
