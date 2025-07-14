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

	*color = ColorANSI(opts, lineType.Title.Colors)
	AddLineFeedBetweenErrorThrown(opts, w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title.Prefix), 1)
}

func FormatTestLine(
	opts Opts,
	line []byte,
	errorBefore *bool,
	stdout io.Writer,
	lineSummary *LineSummary,
) []byte {
	var formattedLine []byte

	if len(line) > 0 {
		line = bytes.TrimRight(line, "\n")
		line = bytes.TrimLeft(line, " ")

		var color []byte

		// manage color and style line depending on the content
		// === RUN
		if bytes.Contains(line, DefaultTitle.Run) {
			if opts.Run.Title.Hide {
				return []byte("")
			}

			if !opts.Run.Summary.Hide {
				lineSummary.Run++
			}

			HandleLineType(opts, &line, opts.Run, DefaultTitle.Run, &color, stdout, errorBefore, false)

			// --- FAIL:
		} else if bytes.Contains(line, DefaultTitle.Fail) {
			if opts.Fail.Title.Hide {
				return []byte("")
			}

			if !opts.Fail.Summary.Hide {
				lineSummary.Fail++
			}

			HandleLineType(opts, &line, opts.Fail, DefaultTitle.Fail, &color, stdout, errorBefore, false)

			// --- PASS:
		} else if bytes.Contains(line, DefaultTitle.Pass) {
			if opts.Pass.Title.Hide {
				return []byte("")
			}

			if !opts.Pass.Summary.Hide {
				lineSummary.Pass++
			}

			HandleLineType(opts, &line, opts.Pass, DefaultTitle.Pass, &color, stdout, errorBefore, false)

			// --- SKIP:
		} else if bytes.Contains(line, DefaultTitle.Skip) {
			if opts.Skip.Title.Hide {
				return []byte("")
			}

			if !opts.Skip.Summary.Hide {
				lineSummary.Skip++
			}

			HandleLineType(opts, &line, opts.Skip, DefaultTitle.Skip, &color, stdout, errorBefore, false)

			// FAIL
		} else if bytes.Equal(line, DefaultTitle.Failed) {
			if opts.Failed.Title.Hide {
				return []byte("")
			}

			if !opts.Failed.Summary.Hide {
				lineSummary.Failed++
			}

			HandleLineType(opts, &line, opts.Failed, DefaultTitle.Failed, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			// ok
		} else if bytes.Equal(line, DefaultTitle.Ok) {
			if opts.Ok.Title.Hide {
				return []byte("")
			}

			if !opts.Ok.Summary.Hide {
				lineSummary.Ok++
			}

			HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			// error thrown
		} else {
			if opts.ErrorThrown.Title.Hide {
				return []byte("")
			}

			if !opts.ErrorThrown.Summary.Hide {
				lineSummary.ErrorThrown++
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

func AddPrintLineSummary(print *[]byte, opts Opts, summary LineTypeSummary, value int) {
	if len(*print) == 0 {
		*print = append(*print, []byte("\n")...)
	}

	*print = append(*print, ColorANSI(opts, summary.Colors)...)
	*print = append(*print, []byte(summary.Prefix)...)
	*print = append(*print, []byte(strconv.Itoa(value))...)
	*print = append(*print, ColorReset...)
	*print = append(*print, []byte("\n")...)
}

func PrintLineSummary(opts Opts, lineSummary LineSummary) []byte {
	var print []byte

	if !opts.Run.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Run.Summary, lineSummary.Run)
	}

	if !opts.Fail.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Fail.Summary, lineSummary.Fail)
	}

	if !opts.Pass.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Pass.Summary, lineSummary.Pass)
	}

	if !opts.Skip.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Skip.Summary, lineSummary.Skip)
	}

	if !opts.Failed.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Failed.Summary, lineSummary.Failed)
	}

	if !opts.Ok.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.Ok.Summary, lineSummary.Ok)
	}

	if !opts.ErrorThrown.Summary.Hide {
		AddPrintLineSummary(&print, opts, opts.ErrorThrown.Summary, lineSummary.ErrorThrown)
	}

	if len(print) > 0 {
		print = append(print, []byte("\n")...)
	}

	return print
}
