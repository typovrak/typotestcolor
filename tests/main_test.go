package tests

import (
	"os"
	"testing"
)

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
