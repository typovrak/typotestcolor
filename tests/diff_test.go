package tests

import (
	"testing"

	"github.com/typovrak/typotestcolor"
)

func TestDiff(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := "test 1"
		got := "test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("1", func(t *testing.T) {
		expected := "test 1"
		got := "test 2"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("2", func(t *testing.T) {
		expected := "test 2"
		got := "test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("3", func(t *testing.T) {
		expected := "test"
		got := "test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("4", func(t *testing.T) {
		expected := "test 1"
		got := "test"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("5", func(t *testing.T) {
		expected := "est 1"
		got := "test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("6", func(t *testing.T) {
		expected := "test 1"
		got := "Test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("7", func(t *testing.T) {
		expected := "  test 1"
		got := "test 1"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})

	t.Run("4", func(t *testing.T) {
		expected := "xx test"
		got := "test"
		typotestcolor.TestDiffStringDefault(t, expected, got)
	})
}
