package tests

import (
	"testing"

	"github.com/typovrak/typotestcolor"
)

type customStringer struct {
	Name string
}

func (c customStringer) String() string {
	return c.Name
}

func TestDiff(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := "test 1"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("1", func(t *testing.T) {
		expected := "test 1"
		got := "test 2"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("2", func(t *testing.T) {
		expected := "test 2"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("3", func(t *testing.T) {
		expected := "test"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("4", func(t *testing.T) {
		expected := "test 1"
		got := "test"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("5", func(t *testing.T) {
		expected := "est 1"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("6", func(t *testing.T) {
		expected := "test 1"
		got := "Test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("7", func(t *testing.T) {
		expected := "  test 1"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("8", func(t *testing.T) {
		expected := "xx test"
		got := "test"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("9", func(t *testing.T) {
		expected := 1
		got := 1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("10", func(t *testing.T) {
		expected := 1
		got := 2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("11", func(t *testing.T) {
		expected := 1.1
		got := 1.1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("12", func(t *testing.T) {
		expected := 1.1
		got := 1.2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("13", func(t *testing.T) {
		expected := []byte("test")
		got := []byte("test")
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("14", func(t *testing.T) {
		expected := []byte("test 1")
		got := []byte("test 2")
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("15", func(t *testing.T) {
		expected := true
		got := true
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("16", func(t *testing.T) {
		expected := false
		got := true
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("17", func(t *testing.T) {
		expected := "Stringer"
		got := customStringer{
			Name: "Stringer",
		}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("18", func(t *testing.T) {
		expected := "Stringer1"
		got := customStringer{
			Name: "Stringer2",
		}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("19", func(t *testing.T) {
		var expected []byte = nil
		var got []byte = nil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("20", func(t *testing.T) {
		var expected []byte = []byte("nil")
		var got []byte = nil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("21", func(t *testing.T) {
		expected := []int{1, 2, 3}
		got := []int{1, 2, 3}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("22", func(t *testing.T) {
		expected := []int{1, 2, 3}
		got := []int{1, 2, 3, 4}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("23", func(t *testing.T) {
		expected := []int{1, 2, 3}
		got := []int{1, 3, 3}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("24", func(t *testing.T) {
		expected := struct{ A int }{A: 1}
		got := struct{ A int }{A: 1}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("25", func(t *testing.T) {
		expected := struct{ A int }{A: 1}
		got := struct{ B int }{B: 1}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("26", func(t *testing.T) {
		expected := struct{ A int }{A: 1}
		got := struct{ A int }{A: 2}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("27", func(t *testing.T) {
		expected := map[string]int{"a": 1}
		got := map[string]int{"a": 1}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("28", func(t *testing.T) {
		expected := map[string]int{"a": 1}
		got := map[string]int{"b": 1}
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("29", func(t *testing.T) {
		expected := map[string]int{"a": 1}
		got := map[string]int{"a": 2}
		typotestcolor.TestDiffDefault(t, expected, got)
	})
}
