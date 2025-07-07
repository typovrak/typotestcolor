package typotestcolor

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

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

	defaultTitle := struct {
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

	// read line by line
	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte

			// manage color and style line depending on the content
			// === RUN
			if bytes.Contains(line, defaultTitle.Run) {
				HandleLineType(&line, opts.Run, defaultTitle.Run, &color, stdout, &errorBefore, false)

				// --- FAIL:
			} else if bytes.Contains(line, defaultTitle.Fail) {
				HandleLineType(&line, opts.Fail, defaultTitle.Fail, &color, stdout, &errorBefore, false)

				// --- PASS:
			} else if bytes.Contains(line, defaultTitle.Pass) {
				HandleLineType(&line, opts.Pass, defaultTitle.Pass, &color, stdout, &errorBefore, false)

				// --- SKIP:
			} else if bytes.Contains(line, defaultTitle.Skip) {
				HandleLineType(&line, opts.Skip, defaultTitle.Skip, &color, stdout, &errorBefore, false)

				// FAIL
			} else if bytes.Equal(line, defaultTitle.Failed) {
				HandleLineType(&line, opts.Failed, defaultTitle.Failed, &color, stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, defaultTitle.Ok) {
				HandleLineType(&line, opts.Ok, defaultTitle.Ok, &color, stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				HandleLineType(&line, opts.ErrorThrown, defaultTitle.ErrorThrown, &color, stdout, &errorBefore, true)
			}

			stdout.Write(color)
			stdout.Write(line)
			stdout.Write([]byte("\033[0m"))
			stdout.Write([]byte("\n"))
		}

		// nothing to read
		if err != nil {
			break
		}
	}

	// restore outputs
	os.Stdout = stdout
	os.Stderr = stderr

	// [0, 125]
	return exitCode
}
