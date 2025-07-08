package tests

import (
	"strconv"
	"testing"

	"github.com/typovrak/typotestcolor"
)

type verifyDefault struct {
	exitCode int
	out      string
}

func validateTestDefault(t *testing.T, res verifyDefault, expected verifyDefault) {
	if res.exitCode != expected.exitCode {
		t.Errorf("[exitCode] expected %d, got %d", expected.exitCode, res.exitCode)
	}

	res.out = strconv.QuoteToASCII(res.out)
	expected.out = strconv.QuoteToASCII(expected.out)

	if res.out != expected.out {
		t.Errorf("[out] expected %s (length: %d), got %s (length: %d)", expected.out, len(expected.out), res.out, len(res.out))
	}
}

func TestDefault(t *testing.T) {
	t.Run("untestable", func(t *testing.T) {
		res := verifyDefault{
			exitCode: 0,
			out:      "",
		}
		expected := verifyDefault{
			exitCode: 0,
			out:      "",
		}

		res.exitCode = typotestcolor.Default(nil)
		validateTestDefault(t, res, expected)
	})
}
