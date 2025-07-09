package tests

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/typovrak/typotestcolor"
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

		typotestcolor.AddLineFeedBetweenErrorThrown(DefaultTestOpts, &res.buffer, &res.errorBefore, isError)
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

		typotestcolor.AddLineFeedBetweenErrorThrown(DefaultTestOpts, &res.buffer, &res.errorBefore, isError)
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

		typotestcolor.AddLineFeedBetweenErrorThrown(DefaultTestOpts, &res.buffer, &res.errorBefore, isError)
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

		typotestcolor.AddLineFeedBetweenErrorThrown(DefaultTestOpts, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Run.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Fail.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Pass.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Skip.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Failed.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().Ok.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
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
			color:       typotestcolor.ColorANSI(DefaultTestOpts, typotestcolor.NewDefaultOpts().ErrorThrown.Colors),
			errorBefore: true,
		}

		typotestcolor.HandleLineType(DefaultTestOpts, &res.line, lineType, defaultLineType, &res.color, &res.buffer, &res.errorBefore, isError)
		validateTestHandleLineType(t, res, expected)
	})
}

type verifyReadTestLines struct {
	out         string
	errorBefore bool
}

func validateTestReadTestLines(t *testing.T, res verifyReadTestLines, expected verifyReadTestLines) {
	resLine := strconv.QuoteToASCII(string(res.out))
	expectedLine := strconv.QuoteToASCII(string(expected.out))

	if resLine != expectedLine {
		t.Errorf("[out] expected %s (length: %d)\n\ngot %s (length: %d)", expectedLine, len(expectedLine), resLine, len(resLine))
	}

	if res.errorBefore != expected.errorBefore {
		t.Errorf("[errorBefore] expected %t, got %t", expected.errorBefore, res.errorBefore)
	}
}

func TestReadTestLines(t *testing.T) {
	t.Run("all line type with ok result", func(t *testing.T) {
		defaultOpts := typotestcolor.NewDefaultOpts()

		res := verifyReadTestLines{
			errorBefore: false,
		}
		expected := verifyReadTestLines{
			errorBefore: false,
		}

		rawOutput := `=== RUN   TestColorANSI
=== RUN   TestColorANSI/run_default_color
html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)
=== RUN   TestColorANSI/fail_default_color
--- PASS: TestColorANSI (0.00s)
    --- PASS: TestColorANSI/run_default_color (0.00s)
    --- PASS: TestColorANSI/fail_default_color (0.00s)
    --- PASS: TestColorANSI/pass_default_color (0.00s)
    --- SKIP: TestColorANSI/skip_default_color (0.00s)
    --- PASS: TestColorANSI/failed_default_color (0.00s)
    --- FAIL: TestColorANSI/ok_default_color (0.00s)
    --- PASS: TestColorANSI/error_thrown_default_color (0.00s)
=== RUN   TestNewDefaultOpts
=== RUN   TestNewDefaultOpts/run_default_style
--- PASS: TestNewDefaultOpts (0.00s)
PASS`

		// === RUN   TestColorANSI
		expected.out = string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestColorANSI/run_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/run_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// \nhtml-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\n
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.ErrorThrown.Colors)) +
			string(defaultOpts.ErrorThrown.Title) +
			"html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)" +
			string(typotestcolor.ColorReset) + "\n\n" +
			// === RUN   TestColorANSI/fail_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/fail_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// --- PASS: TestColorANSI (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/run_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/run_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/fail_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/fail_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/pass_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/pass_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- SKIP: TestColorANSI/skip_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Skip.Colors)) +
			string(defaultOpts.Skip.Title) +
			" TestColorANSI/skip_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/failed_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/failed_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- FAIL: TestColorANSI/ok_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Fail.Colors)) +
			string(defaultOpts.Fail.Title) +
			" TestColorANSI/ok_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/error_thrown_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/error_thrown_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestNewDefaultOpts
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestNewDefaultOpts" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestNewDefaultOpts/run_default_style
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestNewDefaultOpts/run_default_style" +
			string(typotestcolor.ColorReset) + "\n" +
			// --- PASS: TestNewDefaultOpts (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestNewDefaultOpts (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// PASS
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Ok.Colors)) +
			string(defaultOpts.Ok.Title) +
			string(typotestcolor.ColorReset) + "\n"

		var (
			input = bufio.NewReader(strings.NewReader(rawOutput))
			out   bytes.Buffer
		)

		typotestcolor.ReadTestLines(defaultOpts, input, &out, &res.errorBefore)
		res.out = out.String()

		validateTestReadTestLines(t, res, expected)
	})

	t.Run("all line type with failed result", func(t *testing.T) {
		defaultOpts := typotestcolor.NewDefaultOpts()

		res := verifyReadTestLines{
			errorBefore: false,
		}
		expected := verifyReadTestLines{
			errorBefore: false,
		}

		rawOutput := `=== RUN   TestColorANSI
=== RUN   TestColorANSI/run_default_color
html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)
=== RUN   TestColorANSI/fail_default_color
--- PASS: TestColorANSI (0.00s)
    --- PASS: TestColorANSI/run_default_color (0.00s)
    --- PASS: TestColorANSI/fail_default_color (0.00s)
    --- PASS: TestColorANSI/pass_default_color (0.00s)
    --- SKIP: TestColorANSI/skip_default_color (0.00s)
    --- PASS: TestColorANSI/failed_default_color (0.00s)
    --- FAIL: TestColorANSI/ok_default_color (0.00s)
    --- PASS: TestColorANSI/error_thrown_default_color (0.00s)
=== RUN   TestNewDefaultOpts
=== RUN   TestNewDefaultOpts/run_default_style
--- PASS: TestNewDefaultOpts (0.00s)
FAIL`

		// === RUN   TestColorANSI
		expected.out = string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestColorANSI/run_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/run_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// \nhtml-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\n
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.ErrorThrown.Colors)) +
			string(defaultOpts.ErrorThrown.Title) +
			"html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)" +
			string(typotestcolor.ColorReset) + "\n\n" +
			// === RUN   TestColorANSI/fail_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/fail_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// --- PASS: TestColorANSI (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/run_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/run_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/fail_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/fail_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/pass_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/pass_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- SKIP: TestColorANSI/skip_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Skip.Colors)) +
			string(defaultOpts.Skip.Title) +
			" TestColorANSI/skip_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/failed_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/failed_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- FAIL: TestColorANSI/ok_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Fail.Colors)) +
			string(defaultOpts.Fail.Title) +
			" TestColorANSI/ok_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// \t--- PASS: TestColorANSI/error_thrown_default_color (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestColorANSI/error_thrown_default_color (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestNewDefaultOpts
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestNewDefaultOpts" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestNewDefaultOpts/run_default_style
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestNewDefaultOpts/run_default_style" +
			string(typotestcolor.ColorReset) + "\n" +
			// --- PASS: TestNewDefaultOpts (0.00s)
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Pass.Colors)) +
			string(defaultOpts.Pass.Title) +
			" TestNewDefaultOpts (0.00s)" +
			string(typotestcolor.ColorReset) + "\n" +
			// FAIL
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Failed.Colors)) +
			string(defaultOpts.Failed.Title) +
			string(typotestcolor.ColorReset) + "\n"

		var (
			input = bufio.NewReader(strings.NewReader(rawOutput))
			out   bytes.Buffer
		)

		typotestcolor.ReadTestLines(defaultOpts, input, &out, &res.errorBefore)
		res.out = out.String()

		validateTestReadTestLines(t, res, expected)
	})

	t.Run("error thrown with line feed", func(t *testing.T) {
		defaultOpts := typotestcolor.NewDefaultOpts()

		res := verifyReadTestLines{
			errorBefore: false,
		}
		expected := verifyReadTestLines{
			errorBefore: false,
		}

		rawOutput := `=== RUN   TestColorANSI
=== RUN   TestColorANSI/run_default_color
html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\ntest\ntest
=== RUN   TestColorANSI/fail_default_color`

		// === RUN   TestColorANSI
		expected.out = string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestColorANSI/run_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/run_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// \nhtml-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\n
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.ErrorThrown.Colors)) +
			string(defaultOpts.ErrorThrown.Title) +
			"html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)" + "\\n" +
			"test" + "\\n" +
			"test" +
			string(typotestcolor.ColorReset) + "\n\n" +
			// === RUN   TestColorANSI/fail_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/fail_default_color" +
			string(typotestcolor.ColorReset) + "\n"

		var (
			input = bufio.NewReader(strings.NewReader(rawOutput))
			out   bytes.Buffer
		)

		typotestcolor.ReadTestLines(defaultOpts, input, &out, &res.errorBefore)
		res.out = out.String()

		validateTestReadTestLines(t, res, expected)
	})

	t.Run("errorBefore true at the end", func(t *testing.T) {
		defaultOpts := typotestcolor.NewDefaultOpts()

		res := verifyReadTestLines{
			errorBefore: false,
		}
		expected := verifyReadTestLines{
			errorBefore: true,
		}

		rawOutput := `=== RUN   TestColorANSI
=== RUN   TestColorANSI/run_default_color
html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\ntest\ntest`

		// === RUN   TestColorANSI
		expected.out = string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI" +
			string(typotestcolor.ColorReset) + "\n" +
			// === RUN   TestColorANSI/run_default_color
			string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.Run.Colors)) +
			string(defaultOpts.Run.Title) +
			" TestColorANSI/run_default_color" +
			string(typotestcolor.ColorReset) + "\n" +
			// \nhtml-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)\n
			"\n" + string(typotestcolor.ColorANSI(DefaultTestOpts, defaultOpts.ErrorThrown.Colors)) +
			string(defaultOpts.ErrorThrown.Title) +
			"html-minifier_test.go:19: expected <a>  <  </a> (length: 12), got <a>  <</a> (length: 10)" + "\\n" +
			"test" + "\\n" +
			"test" +
			string(typotestcolor.ColorReset) + "\n"

		var (
			input = bufio.NewReader(strings.NewReader(rawOutput))
			out   bytes.Buffer
		)

		typotestcolor.ReadTestLines(defaultOpts, input, &out, &res.errorBefore)
		res.out = out.String()

		validateTestReadTestLines(t, res, expected)
	})
}
