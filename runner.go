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
	Run:         []byte("=== RUN   "),
	Fail:        []byte("--- FAIL: "),
	Pass:        []byte("--- PASS: "),
	Skip:        []byte("--- SKIP: "),
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
	oldlog := log.Writer()

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	os.Stderr = w
	log.SetOutput(w)

	// no error when no test executed
	exitCode := 0

	done := make(chan struct{})
	go func() {
		// INFO: pointers are safe to use unless you declare a nil pointer
		// var print []byte -> *print is always safe
		// var print *[]byte -> print is a nil pointer, this will panic the program
		// so pointer are safe unless you declare it as a pointer already.

		scanner := bufio.NewScanner(r)
		lineTypeBefore := LineTypeEnumNone
		var lineSummary LineSummary
		var aggregationCount AggregationCount

		for scanner.Scan() {
			line := scanner.Bytes()
			formattedLine, aggregationLines := FormatTestLine(opts, line, &lineTypeBefore, &lineSummary, &aggregationCount)

			if len(aggregationLines) > 0 {
				fmt.Fprintf(stdout, string(aggregationLines))
			}

			if len(formattedLine) > 0 {
				fmt.Fprint(stdout, string(formattedLine))
			}
		}

		if aggregationCount.Type != LineTypeEnumNone {
			var aggregationLines []byte

			optsTypeTitleAggregation := GetOptsTypeTitleAggregationFromAggregationCountType(opts, &aggregationCount)
			PrintAggregation(optsTypeTitleAggregation, &aggregationCount, &aggregationLines)

			if len(aggregationLines) > 0 {
				fmt.Fprintf(stdout, string(aggregationLines))
			}
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
