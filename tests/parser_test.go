package tests

import (
	"bytes"
	"strconv"
	"testing"
	"typotestcolor"
)

type verifyAddLineFeedBetweenErrorThrown struct {
	buffer      bytes.Buffer
	errorBefore bool
}

func validateTestAddLineFeedBetweenErrorThrown(t *testing.T, res verifyAddLineFeedBetweenErrorThrown, expected verifyAddLineFeedBetweenErrorThrown) {
	resBuffer := strconv.QuoteToASCII(res.buffer.String())
	expectedBuffer := strconv.QuoteToASCII(expected.buffer.String())

	if resBuffer != expectedBuffer {
		t.Errorf("[buffer] expected %s (length: %d), got %s (length: %d)", expectedBuffer, len(expectedBuffer), resBuffer, len(resBuffer))
	}

	if res.errorBefore != expected.errorBefore {
		t.Errorf("[errorBefore] expected %t, got %t", expected.errorBefore, res.errorBefore)
	}
}

func TestAddLineFeedBetweenErrorThrown(t *testing.T) {
	t.Run("errorBefore and isError should not add \n", func(t *testing.T) {
		isError := true
		res := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: true,
		}
		expected := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: true,
		}

		typotestcolor.AddLineFeedBetweenErrorThrown(&res.buffer, &res.errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, res, expected)
	})

	t.Run("!errorBefore and isError should return \n", func(t *testing.T) {
		isError := true
		res := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: false,
		}
		expected := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: true,
		}
		expected.buffer.Write([]byte("\n"))

		typotestcolor.AddLineFeedBetweenErrorThrown(&res.buffer, &res.errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, res, expected)
	})

	t.Run("errorBefore and !isError should return \n", func(t *testing.T) {
		isError := false
		res := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: true,
		}
		expected := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: false,
		}
		expected.buffer.Write([]byte("\n"))

		typotestcolor.AddLineFeedBetweenErrorThrown(&res.buffer, &res.errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, res, expected)
	})

	t.Run("!errorBefore and !isError should not return \n", func(t *testing.T) {
		isError := false
		res := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: false,
		}
		expected := verifyAddLineFeedBetweenErrorThrown{
			errorBefore: false,
		}

		typotestcolor.AddLineFeedBetweenErrorThrown(&res.buffer, &res.errorBefore, isError)
		validateTestAddLineFeedBetweenErrorThrown(t, res, expected)
	})
}

type verifyHandleLineType struct {
	line        []byte
	color       []byte
	buffer      bytes.Buffer
	errorBefore bool
}

func validateTestHandleLineType(
	t *testing.T,
	res verifyHandleLineType,
	expected verifyHandleLineType,
) {
	resLine := strconv.QuoteToASCII(string(res.line))
	expectedLine := strconv.QuoteToASCII(string(expected.line))

	if resLine != expectedLine {
		t.Errorf("[line] expected %s (length: %d), got %s (length: %d)", expectedLine, len(expectedLine), resLine, len(resLine))
	}

	resColor := strconv.QuoteToASCII(string(res.color))
	expectedColor := strconv.QuoteToASCII(string(expected.color))

	if resColor != expectedColor {
		t.Errorf("[color] expected %s (length: %d), got %s (length: %d)", expectedColor, len(expectedColor), resColor, len(resColor))
	}

	resBuffer := strconv.QuoteToASCII(res.buffer.String())
	expectedBuffer := strconv.QuoteToASCII(expected.buffer.String())

	if resBuffer != expectedBuffer {
		t.Errorf("[buffer] expected %s (length: %d), got %s (length: %d)", expectedBuffer, len(expectedBuffer), resBuffer, len(resBuffer))
	}

	if res.errorBefore != expected.errorBefore {
		t.Errorf("[errorBefore] expected %t, got %t", expected.errorBefore, res.errorBefore)
	}
}

func TestHandleLineType(t *testing.T) {
	t.Run("", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Run
		defaultLineType := typotestcolor.DefaultTitle.Run
		res := verifyHandleLineType{
			line:        []byte(""),
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			line:        []byte(""),
			color:       []byte(""),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})
}
