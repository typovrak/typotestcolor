package typotestcolor

import (
	"bufio"
	"os"
	"testing"
)

var DefaultTitle = struct {
	Run         []byte
	Fail        []byte
	Pass        []byte
	Skip        []byte
	Failed      []byte
	Ok          []byte
	ErrorThrown []byte
}{
	Run:         []byte("=== RUN  "),
	Fail:        []byte("--- FAIL:"),
	Pass:        []byte("--- PASS:"),
	Skip:        []byte("--- SKIP:"),
	Failed:      []byte("FAIL"),
	Ok:          []byte("PASS"),
	ErrorThrown: []byte(""),
}

// return exitCode
func RunTestColor(m *testing.M, opts Opts) int {
	// create a pipe
	r, w, _ := os.Pipe()

	// backup original outputs
	stdout := os.Stdout
	stderr := os.Stderr

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	os.Stderr = w

	exitCode := m.Run()

	// close the writer end of the pipe so the reader stops at EOF
	w.Close()

	// setup the reader
	reader := bufio.NewReader(r)

	errorBefore := false

	// read line by line
	ReadTestLines(opts, reader, stdout, &errorBefore)

	// restore outputs
	os.Stdout = stdout
	os.Stderr = stderr

	// [0, 125]
	return exitCode
}
