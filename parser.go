package typotestcolor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func AddLineFeedBetweenErrorThrown(opts Opts, w io.Writer, errorBefore *bool, isError bool) {
	Debug(opts, "AddLineFeedBetweenErrorThrown")

	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

func HandleLineType(
	opts Opts,
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w io.Writer,
	errorBefore *bool,
	isError bool,
) {
	Debug(opts, "HandleLineType")

	*color = ColorANSI(opts, lineType.Colors)
	AddLineFeedBetweenErrorThrown(opts, w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
}

func ReadTestLines(
	opts Opts,
	reader *bufio.Reader,
	stdout io.Writer,
	errorBefore *bool,
) {
	// TODO: try debugging with log files
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	Debug(opts, "ReadTestLines")

	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			_, err = f.WriteString("\nLEN LINE > 0 :\n" + string(line))
			if err != nil {
				log.Fatal(err)
			}

			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte

			// manage color and style line depending on the content
			// === RUN
			if bytes.Contains(line, DefaultTitle.Run) {
				HandleLineType(opts, &line, opts.Run, DefaultTitle.Run, &color, stdout, errorBefore, false)

				// --- FAIL:
			} else if bytes.Contains(line, DefaultTitle.Fail) {
				HandleLineType(opts, &line, opts.Fail, DefaultTitle.Fail, &color, stdout, errorBefore, false)

				// --- PASS:
			} else if bytes.Contains(line, DefaultTitle.Pass) {
				HandleLineType(opts, &line, opts.Pass, DefaultTitle.Pass, &color, stdout, errorBefore, false)

				// --- SKIP:
			} else if bytes.Contains(line, DefaultTitle.Skip) {
				HandleLineType(opts, &line, opts.Skip, DefaultTitle.Skip, &color, stdout, errorBefore, false)

				// FAIL
			} else if bytes.Equal(line, DefaultTitle.Failed) {
				HandleLineType(opts, &line, opts.Failed, DefaultTitle.Failed, &color, stdout, errorBefore, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, DefaultTitle.Ok) {
				HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				HandleLineType(opts, &line, opts.ErrorThrown, DefaultTitle.ErrorThrown, &color, stdout, errorBefore, true)
			}

			_, err = f.WriteString("\nAFTER SWITCH :\n" + string(line))
			if err != nil {
				log.Fatal(err)
			}

			stdout.Write(color)
			stdout.Write(line)
			stdout.Write(ColorReset)
			stdout.Write([]byte("\n"))

			_, err = f.WriteString("\nAFTER WRITE STDOUT :\n" + string(line))
			if err != nil {
				log.Fatal(err)
			}
		}

		// nothing to read
		if err != nil {

			_, err = f.WriteString("\nERROR :\n" + string(line))
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}
