package tests

import (
	"strconv"
	"testing"

	"github.com/typovrak/typotestcolor"
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
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundCyan],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Run.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("fail default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundRed],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Fail.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("pass default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundGreen],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Pass.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("skip default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundYellow],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Skip.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("failed default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundBlack],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundRed],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Failed.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("ok default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleBold],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundBlack],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundGreen],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.Ok.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})

	t.Run("error thrown default color", func(t *testing.T) {
		res := string(
			typotestcolor.ColorANSI(typotestcolor.ANSIConfig{
				Style:      typotestcolor.ColorANSIStyle[typotestcolor.ANSIStyleNormal],
				Foreground: typotestcolor.ColorANSIForeground[typotestcolor.ANSIForegroundWhite],
				Background: typotestcolor.ColorANSIBackground[typotestcolor.ANSIBackgroundNone],
			}),
		)
		expected := string(typotestcolor.ColorANSI(DefaultTestOpts.ErrorThrown.Title.Colors))
		validateTestColorANSI(t, res, expected)
	})
}
