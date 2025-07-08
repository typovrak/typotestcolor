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
	t.Run("run_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Run
		defaultLineType := typotestcolor.DefaultTitle.Run
		res := verifyHandleLineType{
			// go test run title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default run title
			line: []byte(lineType.Title),
			// default run ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Run.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("fail_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Fail
		defaultLineType := typotestcolor.DefaultTitle.Fail
		res := verifyHandleLineType{
			// go test fail title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default fail title
			line: []byte(lineType.Title),
			// default fail ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Fail.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("pass_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Pass
		defaultLineType := typotestcolor.DefaultTitle.Pass
		res := verifyHandleLineType{
			// go test pass title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default pass title
			line: []byte(lineType.Title),
			// default pass ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Pass.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("skip_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Skip
		defaultLineType := typotestcolor.DefaultTitle.Skip
		res := verifyHandleLineType{
			// go test skip title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default skip title
			line: []byte(lineType.Title),
			// default skip ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Skip.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("failed_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Failed
		defaultLineType := typotestcolor.DefaultTitle.Failed
		res := verifyHandleLineType{
			// go test failed title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default failed title
			line: []byte(lineType.Title),
			// default failed ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Failed.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("ok_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().Ok
		defaultLineType := typotestcolor.DefaultTitle.Ok
		res := verifyHandleLineType{
			// go test ok title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default ok title
			line: []byte(lineType.Title),
			// default ok ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Ok.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})

	t.Run("error_thrown_0", func(t *testing.T) {
		isError := true
		lineType := typotestcolor.NewDefaultOpts().ErrorThrown
		defaultLineType := typotestcolor.DefaultTitle.ErrorThrown
		res := verifyHandleLineType{
			// go test error thrown title
			line:        defaultLineType,
			color:       []byte(""),
			errorBefore: true,
		}
		expected := verifyHandleLineType{
			// default error thrown title
			line: []byte(lineType.Title),
			// default error thrown ansi color
			color:       typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().ErrorThrown.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(&res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})
}
