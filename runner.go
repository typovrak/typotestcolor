package typotestcolor

import (
	"bufio"
	"fmt"
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
	Debug(opts, "RunTestColor")

	// create a pipe
	r, w, _ := os.Pipe()

	// backup original outputs
	stdout := os.Stdout
	stderr := os.Stderr
	oldlog := log.Writer()

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	os.Stderr = w
	log.SetOutput(w)

	// no error when no test executed
	exitCode := 0

	done := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(r)
		errorBefore := false
		var lineSummary LineSummary

		var aggregationCount AggregationCount
		aggregationCount.Type = AggregationTypeNone
		aggregationCount.Value = 0

		for scanner.Scan() {
			line := scanner.Bytes()

			formattedLine := FormatTestLine(opts, line, &errorBefore, w, &lineSummary, aggregationCount)
			fmt.Fprint(stdout, string(formattedLine))
		}

		print := PrintLineSummary(opts, lineSummary)
		if len(print) > 0 {
			fmt.Fprint(stdout, string(print))
		}

		close(done)
	}()

	// test mock
	if m != nil {
		exitCode = m.Run()
	}

	// close the writer end of the pipe so the reader stops at EOF
	w.Close()

	<-done

	// restore outputs
	os.Stdout = stdout
	os.Stderr = stderr
	log.SetOutput(oldlog)

	// [0, 125]
	return exitCode
}
