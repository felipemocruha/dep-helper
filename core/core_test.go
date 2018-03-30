package core

import (
	//	"fmt"
	"testing"
)

func TestCompose(t *testing.T) {
	double := func(s string) string {
		return s + s
	}

	triple := func(s string) string {
		return s + s + s
	}

	expected := "hahahahahaha"
	result := compose(triple, double)("ha")
	if result != expected {
		t.Errorf(
			"Unexpected return value: Want(%v) Have(%v)",
			expected, result)
	}
}

func TestSplitLines(t *testing.T) {
	text := "this\nis\na\ntest"
	expected := []string{"this", "is", "a", "test"}
	result := splitLines(text)

	for i, _ := range result {
		if result[i] != expected[i] {
			t.Errorf(
				"Unexpected return value: Want(%v) Have(%v)",
				expected, result)
		}
	}
}

func TestSplitDep(t *testing.T) {
	text := "gunicorn==19.7.1"
	expected := []string{"gunicorn", "19.7.1"}
	result := splitDep(text)

	for i, _ := range result {
		if result[i] != expected[i] {
			t.Errorf(
				"Unexpected return value: Want(%v) Have(%v)",
				expected, result)
		}
	}
}

func TestParseDepFile(t *testing.T) {
	text := []byte("gunicorn==19.7.1\nnumpy==1.0.0")
	expected := []Dep{
		Dep{"gunicorn", "19.7.1"},
		Dep{"numpy", "1.0.0"}}

	result := ParseDepFile(text)

	for i, _ := range result {
		if result[i] != expected[i] {
			t.Errorf(
				"Unexpected return value: Want(%v) Have(%v)",
				expected, result)
		}
	}
}
