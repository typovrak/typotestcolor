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

func testString1() string {
	return "test 1"
}

func testString2() string {
	return "test 2"
}

func testInt1() int {
	return 1
}

func testInt2() int {
	return 2
}

func testFloat32_1() float32 {
	return 1.1
}

func testFloat32_2() float32 {
	return 1.2
}

func testFloat64_1() float64 {
	return 1.1
}

func testFloat64_2() float64 {
	return 1.2
}

func testBytes1() []byte {
	return []byte("test 1")
}

func testBytes2() []byte {
	return []byte("test 2")
}

func testBool1() bool {
	return true
}

func testBool2() bool {
	return false
}

func testStringer1() customStringer {
	return customStringer{
		Name: "test 1",
	}
}

func testStringer2() customStringer {
	return customStringer{
		Name: "test 2",
	}
}

func testNil() {
}

func testSliceInt1() []int {
	return []int{1, 2, 3}
}

func testSliceInt2() []int {
	return []int{1, 3, 3}
}

func testSliceInt3() []int {
	return []int{1, 2, 3, 4}
}

type testStructA = struct {
	A int
}

type testStructB = struct {
	B int
}

func testStruct1() testStructA {
	return testStructA{
		A: 1,
	}
}

func testStruct2() testStructA {
	return testStructA{
		A: 2,
	}
}

func testStruct3() testStructB {
	return testStructB{
		B: 1,
	}
}

func testMap1() map[string]int {
	return map[string]int{
		"A": 1,
	}
}

func testMap2() map[string]int {
	return map[string]int{
		"A": 2,
	}
}

func testMap3() map[string]int {
	return map[string]int{
		"B": 1,
	}
}

func testAny1() any {
	return nil
}

func testAny2() any {
	return 1
}

func testAny3() any {
	return "test"
}

func TestDiff(t *testing.T) {
	t.Run("strings are equal", func(t *testing.T) {
		expected := testString1()
		got := testString1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("last character is different", func(t *testing.T) {
		expected := testString1()
		got := testString2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("last character is different, reversed", func(t *testing.T) {
		expected := testString2()
		got := testString1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("got string too long", func(t *testing.T) {
		expected := "test"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("got string too short", func(t *testing.T) {
		expected := "test 1"
		got := "test"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("expected string first letter missing", func(t *testing.T) {
		expected := "est 1"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("got string first letter capitalized", func(t *testing.T) {
		expected := "test 1"
		got := "Test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("expected string has trailing left spaces", func(t *testing.T) {
		expected := "  test 1"
		got := "test 1"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("expected string has trailing left letters", func(t *testing.T) {
		expected := "xx test"
		got := "test"
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("ints are equal", func(t *testing.T) {
		expected := testInt1()
		got := testInt1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("ints are differents", func(t *testing.T) {
		expected := testInt1()
		got := testInt2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("float32 are equal", func(t *testing.T) {
		expected := testFloat32_1()
		got := testFloat32_1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("float32 are different", func(t *testing.T) {
		expected := testFloat32_1()
		got := testFloat32_2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("float64 are equal", func(t *testing.T) {
		expected := testFloat64_1()
		got := testFloat64_1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("float64 are different", func(t *testing.T) {
		expected := testFloat64_1()
		got := testFloat64_2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("bytes slices are equal", func(t *testing.T) {
		expected := testBytes1()
		got := testBytes1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("bytes slices are different", func(t *testing.T) {
		expected := testBytes1()
		got := testBytes2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("bools are equal", func(t *testing.T) {
		expected := testBool1()
		got := testBool1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("bools are different", func(t *testing.T) {
		expected := testBool1()
		got := testBool2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("stringer are equal", func(t *testing.T) {
		expected := testStringer1()
		got := testStringer1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("stringer are different", func(t *testing.T) {
		expected := testStringer1()
		got := testStringer2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("nil bytes slices are equal", func(t *testing.T) {
		var expected []byte = nil
		var got []byte = nil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("nil and \"nil\" bytes slices are equal", func(t *testing.T) {
		var expected []byte = []byte("nil")
		var got []byte = nil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("ints slices are equal", func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("ints slices are different, value", func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("ints slices are different, length", func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt3()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("structs are equal", func(t *testing.T) {
		expected := testStruct1()
		got := testStruct1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("structs are different, value", func(t *testing.T) {
		expected := testStruct1()
		got := testStruct2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("structs are different, key", func(t *testing.T) {
		expected := testStruct1()
		got := testStruct3()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("maps are equal", func(t *testing.T) {
		expected := testMap1()
		got := testMap1()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("maps are different, value", func(t *testing.T) {
		expected := testMap1()
		got := testMap2()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("maps are different, key", func(t *testing.T) {
		expected := testMap1()
		got := testMap3()
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	// INFO: testing func() type now
	t.Run("func() strings are equal", func(t *testing.T) {
		expected := testString1
		got := testString1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() last character is different", func(t *testing.T) {
		expected := testString1
		got := testString2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() ints are equal", func(t *testing.T) {
		expected := testInt1
		got := testInt1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() ints are differents", func(t *testing.T) {
		expected := testInt1
		got := testInt2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() float32 are equal", func(t *testing.T) {
		expected := testFloat32_1
		got := testFloat32_1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() float32 are different", func(t *testing.T) {
		expected := testFloat32_1
		got := testFloat32_2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() float64 are equal", func(t *testing.T) {
		expected := testFloat64_1
		got := testFloat64_1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() float64 are different", func(t *testing.T) {
		expected := testFloat64_1
		got := testFloat64_2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() bytes slices are equal", func(t *testing.T) {
		expected := testBytes1
		got := testBytes1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() bytes slices are different", func(t *testing.T) {
		expected := testBytes1
		got := testBytes2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() bools are equal", func(t *testing.T) {
		expected := testBool1
		got := testBool1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() bools are different", func(t *testing.T) {
		expected := testBool1
		got := testBool2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() stringer are equal", func(t *testing.T) {
		expected := testString1
		got := testStringer1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() stringer are different", func(t *testing.T) {
		expected := testString1
		got := testStringer2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() nil are equal", func(t *testing.T) {
		expected := testNil
		got := testNil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() nil and \"nil\" bytes slices are equal", func(t *testing.T) {
		var expected []byte = []byte("nil")
		got := testNil
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() ints slices are equal", func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() ints slices are different, value", func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() ints slices are different, length", func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt3
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() structs are equal", func(t *testing.T) {
		expected := testStruct1
		got := testStruct1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() structs are different, value", func(t *testing.T) {
		expected := testStruct1
		got := testStruct2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() structs are different, key", func(t *testing.T) {
		expected := testStruct1
		got := testStruct3
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() maps are equal", func(t *testing.T) {
		expected := testMap1
		got := testMap1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() maps are different, value", func(t *testing.T) {
		expected := testMap1
		got := testMap2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() maps are different, key", func(t *testing.T) {
		expected := testMap1
		got := testMap3
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() any are equal", func(t *testing.T) {
		expected := testAny1
		got := testAny1
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() any are different, nil and int", func(t *testing.T) {
		expected := testAny1
		got := testAny2
		typotestcolor.TestDiffDefault(t, expected, got)
	})

	t.Run("func() any are different, nil and string", func(t *testing.T) {
		expected := testAny1
		got := testAny3
		typotestcolor.TestDiffDefault(t, expected, got)
	})
}
