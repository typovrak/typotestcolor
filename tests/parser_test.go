package tests

import (
	"bytes"
	"testing"
	"typotestcolor"
)

func validateTestAddLineFeedBetweenErrorThrown(t *testing.T, w string, errorBefore bool, expectedBuffer string, expectedErrorBefore bool) {
	if w != expectedBuffer {
		t.Errorf("expected %s (length: %d), got %s (length: %d)", expectedBuffer, len(expectedBuffer), w, len(w))
	}

	if errorBefore != expectedErrorBefore {
		t.Errorf("expected %t, got %t", errorBefore, expectedErrorBefore)
	}
}

func TestAddLineFeedBetweenErrorThrown(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var (
			w           bytes.Buffer
			errorBefore = true
			isError     = true

			expectedBuffer      = ""
			expectedErrorBefore = true
		)
		typotestcolor.AddLineFeedBetweenErrorThrown(&w, &errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, w.String(), errorBefore, expectedBuffer, expectedErrorBefore)
	})

	t.Run("", func(t *testing.T) {
		var (
			w           bytes.Buffer
			errorBefore = false
			isError     = true

			expectedBuffer      = "\n"
			expectedErrorBefore = true
		)
		typotestcolor.AddLineFeedBetweenErrorThrown(&w, &errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, w.String(), errorBefore, expectedBuffer, expectedErrorBefore)
	})

	t.Run("", func(t *testing.T) {
		var (
			w           bytes.Buffer
			errorBefore = true
			isError     = false

			expectedBuffer      = "\n"
			expectedErrorBefore = false
		)
		typotestcolor.AddLineFeedBetweenErrorThrown(&w, &errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, w.String(), errorBefore, expectedBuffer, expectedErrorBefore)
	})

	t.Run("", func(t *testing.T) {
		var (
			w           bytes.Buffer
			errorBefore = false
			isError     = false

			expectedBuffer      = ""
			expectedErrorBefore = false
		)
		typotestcolor.AddLineFeedBetweenErrorThrown(&w, &errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, w.String(), errorBefore, expectedBuffer, expectedErrorBefore)
	})
}

func TestHandleLineType(t *testing.T) {
}
