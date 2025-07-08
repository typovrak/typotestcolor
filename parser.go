package typotestcolor

import (
	"bufio"
	"bytes"
	"io"
)

func AddLineFeedBetweenErrorThrown(w io.Writer, errorBefore *bool, isError bool) {
	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

func HandleLineType(
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w io.Writer,
	errorBefore *bool,
	isError bool,
) {
	*color = ColorANSI(lineType.Colors)
	AddLineFeedBetweenErrorThrown(w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
}

func ReadTestLines(
	opts Opts,
	reader *bufio.Reader,
	stdout io.Writer,
	errorBefore *bool,
) {
	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte

			// manage color and style line depending on the content
			// === RUN
			if bytes.Contains(line, DefaultTitle.Run) {
				HandleLineType(&line, opts.Run, DefaultTitle.Run, &color, stdout, errorBefore, false)

				// --- FAIL:
			} else if bytes.Contains(line, DefaultTitle.Fail) {
				HandleLineType(&line, opts.Fail, DefaultTitle.Fail, &color, stdout, errorBefore, false)

				// --- PASS:
			} else if bytes.Contains(line, DefaultTitle.Pass) {
				HandleLineType(&line, opts.Pass, DefaultTitle.Pass, &color, stdout, errorBefore, false)

				// --- SKIP:
			} else if bytes.Contains(line, DefaultTitle.Skip) {
				HandleLineType(&line, opts.Skip, DefaultTitle.Skip, &color, stdout, errorBefore, false)

				// FAIL
			} else if bytes.Equal(line, DefaultTitle.Failed) {
				HandleLineType(&line, opts.Failed, DefaultTitle.Failed, &color, stdout, errorBefore, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, DefaultTitle.Ok) {
				HandleLineType(&line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				HandleLineType(&line, opts.ErrorThrown, DefaultTitle.ErrorThrown, &color, stdout, errorBefore, true)
			}

			stdout.Write(color)
			stdout.Write(line)
			stdout.Write(ColorReset)
			stdout.Write([]byte("\n"))
		}

		// nothing to read
		if err != nil {
			break
		}
	}
}
