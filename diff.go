package typotestcolor

import (
	"fmt"
	"strings"
	"testing"
)

func DiffPrintColor(color ANSIForeground, highlight bool) []byte {
	var (
		opts           = NewDefaultOpts()
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
	}

	print = append(print, ColorANSI(opts, ansiConfig)...)
	return print
}

func AddStringInSlice(slice string, i int, string string) string {
	// prevent slice overflow
	if i < 0 || i > len(slice) {
		return slice
	}

	return slice[:i] + string + slice[i:]
}

func TestDiffString(t *testing.T, expected string, got string) {
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

	expectedLen := len(expected)
	gotLen := len(got)

	gotErrorBefore := false

	var gotSlices []string
	lastAddedIndex := 0

	for i := 0; i < gotLen; i++ {
		// diff
		// got[i] and expected[i] are always defined
		if i >= expectedLen || expected[i] != got[i] {

			if !gotErrorBefore {
				gotSlices = append(gotSlices, got[lastAddedIndex:i])
				gotSlices = append(gotSlices, string(DiffPrintColor(ANSIForegroundRed, true)))

				gotErrorBefore = true
				lastAddedIndex = i
			}

			continue
		}

		// same
		if gotErrorBefore {
			gotSlices = append(gotSlices, got[lastAddedIndex:i])
			gotSlices = append(gotSlices, string(ColorReset))
			gotSlices = append(gotSlices, string(DiffPrintColor(ANSIForegroundRed, false)))

			gotErrorBefore = false
			lastAddedIndex = i
		}

	}

	if lastAddedIndex != gotLen {
		gotSlices = append(gotSlices, got[lastAddedIndex:gotLen])
	}

	print.WriteByte('\n')

	// expected part
	print.WriteString(expectedPrefix)
	print.Write(DiffPrintColor(ANSIForegroundGreen, false))
	print.WriteString(expected)
	print.WriteByte('\n')

	// got part
	print.WriteString(gotPrefix)
	print.Write(DiffPrintColor(ANSIForegroundRed, false))
	print.WriteString(strings.Join(gotSlices, ""))
	print.WriteByte('\n')

	t.Error(print.String())
	// t.Error(strconv.QuoteToASCII(print.String()))
}

// mettre expected en full vert
// mettre got en full rouge

// expected: test 2
// got: test \1\

// expected: t 2
// got: t\est\ 2

// expected: t/est/ 2
// got: t 2
