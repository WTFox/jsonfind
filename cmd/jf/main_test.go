package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_splitToParentPath(t *testing.T) {
	is := is.New(t)

	testData := []struct {
		value,
		expected string
	}{
		{".[2]", "."},
		{".[2].eyeColor", ".[2]"},
		{".[2].eyeColor.foobar", ".[2].eyeColor"},
		{".[2].eyeColors[2].foobar", ".[2].eyeColors[2]"},
		{".[2].eyeColors[2].foobar[3]", ".[2].eyeColors[2].foobar"},
	}

	t.Run("splitToParentPath", func(t *testing.T) {
		for _, td := range testData {
			is.Equal(splitToParentPath(td.value), td.expected)
		}
	})
}
