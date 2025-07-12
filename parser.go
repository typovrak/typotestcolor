package typotestcolor

import (
	"bytes"
	"io"
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
	// AddLineFeedBetweenErrorThrown(opts, w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
}

func FormatTestLine(
	opts Opts,
	line []byte,
	errorBefore *bool,
	stdout io.Writer,
) []byte {
	var formattedLine []byte

	if len(line) > 0 {
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
			formattedLine = append(formattedLine, []byte("\n")...)

			// ok
		} else if bytes.Equal(line, DefaultTitle.Ok) {
			HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			// error thrown
		} else {
			HandleLineType(opts, &line, opts.ErrorThrown, DefaultTitle.ErrorThrown, &color, stdout, errorBefore, true)
		}

		if len(color) > 0 {
			formattedLine = append(formattedLine, color...)
		}

		formattedLine = append(formattedLine, line...)
		formattedLine = append(formattedLine, ColorReset...)
		formattedLine = append(formattedLine, []byte("\n")...)
	}

	return formattedLine
}
