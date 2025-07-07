package main

import "testing"

func Default(m *testing.M) int {
	return RunTestColor(m, NewDefaultOpts())
}
