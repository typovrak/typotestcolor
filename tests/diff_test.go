package tests

import (
	"testing"

	"github.com/typovrak/typotestcolor"
)

func TestDiff(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := "test 1"
		got := "test 1"
		typotestcolor.TestDiffString(t, expected, got)
	})

	t.Run("1", func(t *testing.T) {
		expected := "test 1"
		got := "test 2"
		typotestcolor.TestDiffString(t, expected, got)
	})

	t.Run("2", func(t *testing.T) {
		expected := "test 2"
		got := "test 1"
		typotestcolor.TestDiffString(t, expected, got)
	})

	t.Run("3", func(t *testing.T) {
		expected := "test"
		got := "test 1"
		typotestcolor.TestDiffString(t, expected, got)
	})

	t.Run("4", func(t *testing.T) {
		expected := "test 1"
		got := "test"
		typotestcolor.TestDiffString(t, expected, got)
	})
}
