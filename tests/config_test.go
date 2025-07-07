package tests

import (
	"testing"
	"typotestcolor"
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
		res := typotestcolor.NewDefaultOpts().Run.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Run.Colors.Foreground
		expected := 36
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Run.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("run default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Run.Title
		expected := "\t=== RUN:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("fail default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Fail.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Fail.Colors.Foreground
		expected := 31
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Fail.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("fail default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Fail.Title
		expected := "\t--- FAIL:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("pass default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Pass.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Pass.Colors.Foreground
		expected := 32
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Pass.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("pass default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Pass.Title
		expected := "\t--- PASS:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("skip default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Skip.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Skip.Colors.Foreground
		expected := 33
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Skip.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("skip default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Skip.Title
		expected := "\t--- SKIP:"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("failed default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Failed.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Failed.Colors.Foreground
		expected := 30
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Failed.Colors.Background
		expected := 41
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("failed default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Failed.Title
		expected := "FAIL"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("ok default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Ok.Colors.Style
		expected := 1
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Ok.Colors.Foreground
		expected := 30
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Ok.Colors.Background
		expected := 42
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("ok default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().Ok.Title
		expected := "PASS"
		validateTestNewDefaultOptsString(t, res, expected)
	})

	t.Run("error thrown default style", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().ErrorThrown.Colors.Style
		expected := 22
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default foreground", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().ErrorThrown.Colors.Foreground
		expected := 37
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default background", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().ErrorThrown.Colors.Background
		expected := 0
		validateTestNewDefaultOptsInt(t, res, expected)
	})

	t.Run("error thrown default title", func(t *testing.T) {
		res := typotestcolor.NewDefaultOpts().ErrorThrown.Title
		expected := ""
		validateTestNewDefaultOptsString(t, res, expected)
	})
}
