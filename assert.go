package typotestcolor

import "testing"

func AssertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Error("expected an error, got no error")
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected: no error, got: %s", err.Error())
	}
}
