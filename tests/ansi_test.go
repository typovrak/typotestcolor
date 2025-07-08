package tests

import (
	"strconv"
	"testing"
	"typotestcolor"
)

func validateTestColorANSI(t *testing.T, res string, expected string) {
	res = strconv.QuoteToASCII(res)
	expected = strconv.QuoteToASCII(expected)

	if res != expected {
		t.Errorf("expected %s (length: %d), got %s (length: %d)", expected, len(expected), res, len(res))
	}
}

func TestColorANSI(t *testing.T) {
	t.Run("run default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundCyan],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Run.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("fail default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundRed],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Fail.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("pass default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundGreen],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Pass.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("skip default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundYellow],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Skip.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("failed default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundBlack],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundRed],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Failed.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("ok default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundBlack],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundGreen],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().Ok.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("error thrown default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSISTyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundWhite],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(typotestcolor.NewDefaultOpts().ErrorThrown.Colors))
		validateTestColorANSI(t, res, expected)
	})
}
