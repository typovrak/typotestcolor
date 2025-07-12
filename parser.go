package typotestcolor

import (
	"bytes"
	"io"
	"strconv"
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

func FormatTestLine(
	opts Opts,
	line []byte,
	errorBefore *bool,
	stdout io.Writer,
	lineAggregation *LineAggregation,
) []byte {
	var formattedLine []byte

	if len(line) > 0 {
		line = bytes.TrimRight(line, "\n")
		line = bytes.TrimLeft(line, " ")

		var color []byte

		// manage color and style line depending on the content
		// === RUN
		if bytes.Contains(line, DefaultTitle.Run) {
			if opts.Run.Hide {
				return []byte("")
			}

			if !opts.Run.AggregationHide {
				lineAggregation.Run++
			}

			HandleLineType(opts, &line, opts.Run, DefaultTitle.Run, &color, stdout, errorBefore, false)

			// --- FAIL:
		} else if bytes.Contains(line, DefaultTitle.Fail) {
			if opts.Fail.Hide {
				return []byte("")
			}

			if !opts.Fail.AggregationHide {
				lineAggregation.Fail++
			}

			HandleLineType(opts, &line, opts.Fail, DefaultTitle.Fail, &color, stdout, errorBefore, false)

			// --- PASS:
		} else if bytes.Contains(line, DefaultTitle.Pass) {
			if opts.Pass.Hide {
				return []byte("")
			}

			if !opts.Pass.AggregationHide {
				lineAggregation.Pass++
			}

			HandleLineType(opts, &line, opts.Pass, DefaultTitle.Pass, &color, stdout, errorBefore, false)

			// --- SKIP:
		} else if bytes.Contains(line, DefaultTitle.Skip) {
			if opts.Skip.Hide {
				return []byte("")
			}

			if !opts.Skip.AggregationHide {
				lineAggregation.Skip++
			}

			HandleLineType(opts, &line, opts.Skip, DefaultTitle.Skip, &color, stdout, errorBefore, false)

			// FAIL
		} else if bytes.Equal(line, DefaultTitle.Failed) {
			if opts.Failed.Hide {
				return []byte("")
			}

			if !opts.Failed.AggregationHide {
				lineAggregation.Failed++
			}

			HandleLineType(opts, &line, opts.Failed, DefaultTitle.Failed, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			// ok
		} else if bytes.Equal(line, DefaultTitle.Ok) {
			if opts.Ok.Hide {
				return []byte("")
			}

			if !opts.Ok.AggregationHide {
				lineAggregation.Ok++
			}

			HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			// error thrown
		} else {
			if opts.ErrorThrown.Hide {
				return []byte("")
			}

			if !opts.ErrorThrown.AggregationHide {
				lineAggregation.ErrorThrown++
			}

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

func AddPrintLineAggregation(print *[]byte, title string, value int) {
	*print = append(*print, []byte(title)...)
	*print = append(*print, []byte(strconv.Itoa(value))...)
	*print = append(*print, []byte("\n")...)
}

func PrintLineAggregation(opts Opts, lineAggregation LineAggregation) []byte {
	var print []byte

	if !opts.Run.AggregationHide {
		AddPrintLineAggregation(&print, opts.Run.AggregationTitle, lineAggregation.Run)
	}

	if !opts.Fail.AggregationHide {
		AddPrintLineAggregation(&print, opts.Fail.AggregationTitle, lineAggregation.Fail)
	}

	if !opts.Pass.AggregationHide {
		AddPrintLineAggregation(&print, opts.Pass.AggregationTitle, lineAggregation.Pass)
	}

	if !opts.Skip.AggregationHide {
		AddPrintLineAggregation(&print, opts.Skip.AggregationTitle, lineAggregation.Skip)
	}

	if !opts.Failed.AggregationHide {
		AddPrintLineAggregation(&print, opts.Failed.AggregationTitle, lineAggregation.Failed)
	}

	if !opts.Ok.AggregationHide {
		AddPrintLineAggregation(&print, opts.Ok.AggregationTitle, lineAggregation.Ok)
	}

	if !opts.ErrorThrown.AggregationHide {
		AddPrintLineAggregation(&print, opts.ErrorThrown.AggregationTitle, lineAggregation.ErrorThrown)
	}

	return print
}
