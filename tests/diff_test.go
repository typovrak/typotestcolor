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
	index := 0

	t.Run(typotestcolor.RunTitle(&index, "strings are equal"), func(t *testing.T) {
		expected := testString1()
		got := testString1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "last character is different"), func(t *testing.T) {
		expected := testString1()
		got := testString2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "last character is different, reversed"), func(t *testing.T) {
		expected := testString2()
		got := testString1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "got string too long"), func(t *testing.T) {
		expected := "test"
		got := "test 1"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "got string too short"), func(t *testing.T) {
		expected := "test 1"
		got := "test"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "expected string first letter missing"), func(t *testing.T) {
		expected := "est 1"
		got := "test 1"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "got string first letter capitalized"), func(t *testing.T) {
		expected := "test 1"
		got := "Test 1"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "expected string has trailing left spaces"), func(t *testing.T) {
		expected := "  test 1"
		got := "test 1"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "expected string has trailing left letters"), func(t *testing.T) {
		expected := "xx test"
		got := "test"

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "ints are equal"), func(t *testing.T) {
		expected := testInt1()
		got := testInt1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "ints are differents"), func(t *testing.T) {
		expected := testInt1()
		got := testInt2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "float32 are equal"), func(t *testing.T) {
		expected := testFloat32_1()
		got := testFloat32_1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "float32 are different"), func(t *testing.T) {
		expected := testFloat32_1()
		got := testFloat32_2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "float64 are equal"), func(t *testing.T) {
		expected := testFloat64_1()
		got := testFloat64_1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "float64 are different"), func(t *testing.T) {
		expected := testFloat64_1()
		got := testFloat64_2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "bytes slices are equal"), func(t *testing.T) {
		expected := testBytes1()
		got := testBytes1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "bytes slices are different"), func(t *testing.T) {
		expected := testBytes1()
		got := testBytes2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	// TODO: mettre en place un contains pour le assert error afin que cela soit la bonne erreur
	// ou une autre fonction

	// TODO: mettre un paramètre pour que les 2 types doivent être égaux???
	// -> TestDiffStrict()

	t.Run(typotestcolor.RunTitle(&index, "bools are equal"), func(t *testing.T) {
		expected := testBool1()
		got := testBool1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "bools are different"), func(t *testing.T) {
		expected := testBool1()
		got := testBool2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "stringer are equal"), func(t *testing.T) {
		expected := testStringer1()
		got := testStringer1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "stringer are different"), func(t *testing.T) {
		expected := testStringer1()
		got := testStringer2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "nil bytes slices are equal"), func(t *testing.T) {
		var expected []byte = nil
		var got []byte = nil

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "nil and \"nil\" bytes slices are different"), func(t *testing.T) {
		var expected []byte = []byte("nil")
		var got []byte = nil

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "ints slices are equal"), func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "ints slices are different, value"), func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "ints slices are different, length"), func(t *testing.T) {
		expected := testSliceInt1()
		got := testSliceInt3()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "structs are equal"), func(t *testing.T) {
		expected := testStruct1()
		got := testStruct1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "structs are different, value"), func(t *testing.T) {
		expected := testStruct1()
		got := testStruct2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "structs are different, key"), func(t *testing.T) {
		expected := testStruct1()
		got := testStruct3()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "maps are equal"), func(t *testing.T) {
		expected := testMap1()
		got := testMap1()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "maps are different, value"), func(t *testing.T) {
		expected := testMap1()
		got := testMap2()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "maps are different, key"), func(t *testing.T) {
		expected := testMap1()
		got := testMap3()

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	// INFO: testing func() type now
	t.Run(typotestcolor.RunTitle(&index, "func() strings are equal"), func(t *testing.T) {
		expected := testString1
		got := testString1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() last character is different"), func(t *testing.T) {
		expected := testString1
		got := testString2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() ints are equal"), func(t *testing.T) {
		expected := testInt1
		got := testInt1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() ints are differents"), func(t *testing.T) {
		expected := testInt1
		got := testInt2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() float32 are equal"), func(t *testing.T) {
		expected := testFloat32_1
		got := testFloat32_1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() float32 are different"), func(t *testing.T) {
		expected := testFloat32_1
		got := testFloat32_2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() float64 are equal"), func(t *testing.T) {
		expected := testFloat64_1
		got := testFloat64_1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() float64 are different"), func(t *testing.T) {
		expected := testFloat64_1
		got := testFloat64_2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() bytes slices are equal"), func(t *testing.T) {
		expected := testBytes1
		got := testBytes1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() bytes slices are different"), func(t *testing.T) {
		expected := testBytes1
		got := testBytes2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() bools are equal"), func(t *testing.T) {
		expected := testBool1
		got := testBool1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() bools are different"), func(t *testing.T) {
		expected := testBool1
		got := testBool2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() stringer are equal"), func(t *testing.T) {
		expected := testString1
		got := testStringer1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() stringer are different"), func(t *testing.T) {
		expected := testString1
		got := testStringer2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() nil are equal"), func(t *testing.T) {
		expected := testNil
		got := testNil

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() nil and \"nil\" bytes slices are equal"), func(t *testing.T) {
		var expected []byte = []byte("nil")
		got := testNil

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() ints slices are equal"), func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() ints slices are different, value"), func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() ints slices are different, length"), func(t *testing.T) {
		expected := testSliceInt1
		got := testSliceInt3

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() structs are equal"), func(t *testing.T) {
		expected := testStruct1
		got := testStruct1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() structs are different, value"), func(t *testing.T) {
		expected := testStruct1
		got := testStruct2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() structs are different, key"), func(t *testing.T) {
		expected := testStruct1
		got := testStruct3

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() maps are equal"), func(t *testing.T) {
		expected := testMap1
		got := testMap1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() maps are different, value"), func(t *testing.T) {
		expected := testMap1
		got := testMap2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() maps are different, key"), func(t *testing.T) {
		expected := testMap1
		got := testMap3

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() any are equal"), func(t *testing.T) {
		expected := testAny1
		got := testAny1

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertNoError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() any are different, nil and int"), func(t *testing.T) {
		expected := testAny1
		got := testAny2

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})

	t.Run(typotestcolor.RunTitle(&index, "func() any are different, nil and string"), func(t *testing.T) {
		expected := testAny1
		got := testAny3

		err := typotestcolor.TestDiffDefault(expected, got)
		typotestcolor.AssertError(t, err)
	})
}
