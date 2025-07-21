package typotestcolor

import (
	"reflect"
	"strings"
	"testing"
)

func AssertSameType(t *testing.T, expected any, got any) {
	t.Helper()

	expectedType := reflect.TypeOf(expected)
	gotType := reflect.TypeOf(got)

	if expectedType != gotType {
		t.Errorf("expected: %s, got: %s", expectedType.String(), gotType.String())
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Error("expected: an error, got: no error")
	}
}

func AssertErrorStrict(t *testing.T, err error, contains string) {
	t.Helper()
	errMessage := err.Error()

	if err != nil && !strings.Contains(errMessage, contains) {
		t.Errorf("expected: %s, got: %s", contains, errMessage)
		return
	}

	if err == nil {
		t.Error("expected: an error, got: no error")
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected: no error, got: %s", err.Error())
	}
}
