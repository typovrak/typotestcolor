package tests

import (
	"strconv"
	"testing"

	"github.com/typovrak/typotestcolor"
)

type verifyRunTestColor struct {
	exitCode int
	out      string
}

func validateTestRunTestColor(t *testing.T, res verifyRunTestColor, expected verifyRunTestColor) {
	if res.exitCode != expected.exitCode {
		t.Errorf("[exitCode] expected %d, got %d", expected.exitCode, res.exitCode)
	}

	res.out = strconv.QuoteToASCII(res.out)
	expected.out = strconv.QuoteToASCII(expected.out)

	if res.out != expected.out {
		t.Errorf("[out] expected %s (length: %d), got %s (length: %d)", expected.out, len(expected.out), res.out, len(res.out))
	}
}

func TestRunTestColor(t *testing.T) {
	t.Run("untestable", func(t *testing.T) {
		res := verifyRunTestColor{
			exitCode: 0,
			out:      "",
		}
		expected := verifyRunTestColor{
			exitCode: 0,
			out:      "",
		}

		res.exitCode = typotestcolor.RunTestColor(nil, typotestcolor.NewDefaultOpts())
		validateTestRunTestColor(t, res, expected)
	})
}
