package tests

import (
	"os"
	"testing"

	"github.com/typovrak/typotestcolor"
)

var DefaultTestOpts = typotestcolor.NewDefaultOpts()

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	exitCode := typotestcolor.RunTestColor(m, DefaultTestOpts)
	os.Exit(exitCode)
}
