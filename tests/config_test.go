package tests

import (
	"testing"
)

func validateTestNewDefaultOptsInt(t *testing.T, res int, expected int) {
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}

func validateTestNewDefaultOptsString(t *testing.T, res string, expected string) {
	if res != expected {
		t.Errorf("expected %s (length: %d), got %s (length: %d)", expected, len(expected), res, len(res))
	}
}

func TestNewDefaultOpts(t *testing.T) {
	t.Run("run default style", func(t *testing.T) {
		res := DefaultTestOpts.Run.Title.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Run.Title.Colors.Foreground
		expected := 36
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default background", func(t *testing.T) {
		res := DefaultTestOpts.Run.Title.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default title", func(t *testing.T) {
		res := DefaultTestOpts.Run.Title.Prefix
		expected := "\t=== RUN:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("fail default style", func(t *testing.T) {
		res := DefaultTestOpts.Fail.Title.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Fail.Title.Colors.Foreground
		expected := 31
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default background", func(t *testing.T) {
		res := DefaultTestOpts.Fail.Title.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default title", func(t *testing.T) {
		res := DefaultTestOpts.Fail.Title.Prefix
		expected := "\t--- FAIL:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("pass default style", func(t *testing.T) {
		res := DefaultTestOpts.Pass.Title.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Pass.Title.Colors.Foreground
		expected := 32
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default background", func(t *testing.T) {
		res := DefaultTestOpts.Pass.Title.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default title", func(t *testing.T) {
		res := DefaultTestOpts.Pass.Title.Prefix
		expected := "\t--- PASS:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("skip default style", func(t *testing.T) {
		res := DefaultTestOpts.Skip.Title.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Skip.Title.Colors.Foreground
		expected := 33
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default background", func(t *testing.T) {
		res := DefaultTestOpts.Skip.Title.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default title", func(t *testing.T) {
		res := DefaultTestOpts.Skip.Title.Prefix
		expected := "\t--- SKIP:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("failed default style", func(t *testing.T) {
		res := DefaultTestOpts.Failed.Title.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Failed.Title.Colors.Foreground
		expected := 30
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default background", func(t *testing.T) {
		res := DefaultTestOpts.Failed.Title.Colors.Background
		expected := 41
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default title", func(t *testing.T) {
		res := DefaultTestOpts.Failed.Title.Prefix
		expected := "FAIL"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("ok default style", func(t *testing.T) {
		res := DefaultTestOpts.Ok.Title.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default foreground", func(t *testing.T) {
		res := DefaultTestOpts.Ok.Title.Colors.Foreground
		expected := 30
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default background", func(t *testing.T) {
		res := DefaultTestOpts.Ok.Title.Colors.Background
		expected := 42
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default title", func(t *testing.T) {
		res := DefaultTestOpts.Ok.Title.Prefix
		expected := "PASS"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("error thrown default style", func(t *testing.T) {
		res := DefaultTestOpts.ErrorThrown.Title.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default foreground", func(t *testing.T) {
		res := DefaultTestOpts.ErrorThrown.Title.Colors.Foreground
		expected := 37
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default background", func(t *testing.T) {
		res := DefaultTestOpts.ErrorThrown.Title.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default title", func(t *testing.T) {
		res := DefaultTestOpts.ErrorThrown.Title.Prefix
		expected := ""
		validateTestNewDefaultOptsString(t, res, expected)
	})
}
