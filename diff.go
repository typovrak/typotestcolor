package typotestcolor

import (
	"fmt"
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

func TestDiffString(t *testing.T, expected string, got string) {
	if expected == got {
		return
	}

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

	//for i := 0; i < min(expectedLen, gotLen); i++ {
	//	if expected[i] != got[i] {
	//		// diff
	//	}
	//}

	t.Errorf(
		"\n%s%s%s%s\n%s%s%s%s",
		expectedPrefix,
		DiffPrintColor(ANSIForegroundGreen, false),
		expected,
		ColorReset,
		gotPrefix,
		DiffPrintColor(ANSIForegroundRed, false),
		got,
		ColorReset,
	)
}

// mettre expected en full vert
// mettre got en full rouge

// expected: test 2
// got: test \1\

// expected: t 2
// got: t\est\ 2

// expected: t/est/ 2
// got: t 2
