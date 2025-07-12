package typotestcolor

import (
	"bufio"
	"log"
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
	// TODO: try debugging with log files
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("ENTER IN RunTestColor")

	Debug(opts, "RunTestColor")

	// create a pipe
	r, w, _ := os.Pipe()

	f.WriteString("create a pipe")

	// backup original outputs
	stdout := os.Stdout
	// stderr := os.Stderr

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	// os.Stderr = w

	// no error when no test executed
	exitCode := 0

	f.WriteString("BEFORE RUNNING TEST")

	// test mock
	if m != nil {
		exitCode = m.Run()
	}

	f.WriteString("AFTER RUNNING TEST")
	// close the writer end of the pipe so the reader stops at EOF
	w.Close()

	// setup the reader
	reader := bufio.NewReader(r)

	f.WriteString("AFTER READER BUFIO")

	errorBefore := false

	// read line by line
	ReadTestLines(opts, reader, stdout, &errorBefore)

	f.WriteString("AFTER ReadTestLines")

	// restore outputs
	os.Stdout = stdout
	// os.Stderr = stderr

	f.WriteString("BEFORE RETURNING EXITCODE")

	// [0, 125]
	return exitCode
}
