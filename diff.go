package typotestcolor

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func DiffPrintColor(color ANSIForeground, highlight bool, opts Opts) []byte {
	var (
		print          []byte
		ansiConfig     ANSIConfig
		ansiForeground ANSIForeground
		ansiBackground ANSIBackground
	)

	switch color {
	case ANSIForegroundGreen:
		ansiForeground = ANSIForegroundGreen
		ansiBackground = ANSIBackgroundGreen

	case ANSIForegroundRed:
		ansiForeground = ANSIForegroundRed
		ansiBackground = ANSIBackgroundRed

	default:
		ansiForeground = ANSIForegroundBlue
		ansiBackground = ANSIBackgroundBlue
	}

	if highlight {
		ansiConfig = ANSIConfig{
			Style:      ColorANSIStyle[ANSIStyleNormal],
			Foreground: ColorANSIForeground[ANSIForegroundBlack],
			Background: ColorANSIBackground[ansiBackground],
		}
	} else {
		ansiConfig = ANSIConfig{
			Style:      ColorANSIStyle[ANSIStyleNormal],
			Foreground: ColorANSIForeground[ansiForeground],
			Background: ColorANSIBackground[ANSIBackgroundNone],
		}

		// INFO: need to reset for ANSIBackgroundNone to works
		print = append(print, ColorReset...)
	}

	print = append(print, ColorANSI(opts, ansiConfig)...)
	return print
}

type TestDiffOpts = struct {
	opts         Opts
	printToASCII bool
}

func TestDiffNewDefaultOpts() TestDiffOpts {
	return TestDiffOpts{
		opts:         NewDefaultOpts(),
		printToASCII: false,
	}
}

func TestDiffStringDefault(t *testing.T, expected string, got string) {
	TestDiffString(t, expected, got, TestDiffNewDefaultOpts())
}

func TestDiffString(t *testing.T, expected string, got string, opts TestDiffOpts) {
	if expected == got {
		return
	}

	var print strings.Builder

	expectedPrefix := fmt.Sprintf("expected (length: %d): ", len(expected))
	gotPrefix := fmt.Sprintf("got (length: %d): ", len(got))

	expectedPrefixLen := len(expectedPrefix)
	gotPrefixLen := len(gotPrefix)

	if expectedPrefixLen > gotPrefixLen {
		for i := 0; i < expectedPrefixLen-gotPrefixLen; i++ {
			gotPrefix += " "
		}
	} else if gotPrefixLen > expectedPrefixLen {
		for i := 0; i < gotPrefixLen-expectedPrefixLen; i++ {
			expectedPrefix += " "
		}
	}

	// start header print
	print.WriteByte('\n')

	// expected part
	print.WriteString(expectedPrefix)
	print.Write(DiffPrintColor(ANSIForegroundGreen, false, opts.opts))
	print.WriteString(expected)
	print.WriteByte('\n')

	// got part
	print.WriteString(gotPrefix)
	print.Write(DiffPrintColor(ANSIForegroundRed, false, opts.opts))

	expectedLen := len(expected)
	gotLen := len(got)

	errorBefore := false

	// build got print
	for i := 0; i < min(expectedLen, gotLen); i++ {
		// same
		if expected[i] == got[i] {
			// remove error highlight
			if errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, false, opts.opts))
				errorBefore = false
			}

			// diff
		} else {
			// add error highlight
			if !errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, true, opts.opts))
				errorBefore = true
			}
		}

		print.WriteByte(got[i])
	}

	// reset ANSI styles
	print.Write(ColorReset)
	print.WriteByte('\n')

	if opts.printToASCII {
		t.Error(strconv.QuoteToASCII(print.String()))
	} else {
		t.Error(print.String())
	}
}
