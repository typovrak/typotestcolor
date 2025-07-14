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
	// *line = append(*line, []byte(lineType.Title.Suffix)...)
}

func PrintAggregation(opts Opts, optsTypeTitleAggregation LineTypeTitleAggregation, aggregationCount *AggregationCount, aggregationLines *[]byte) {
	if aggregationCount.Value >= optsTypeTitleAggregation.Threshold && len(aggregationCount.Lines) > 2 {
		// first line
		*aggregationLines = append(*aggregationLines, aggregationCount.Lines[0]...)

		// setup aggregation color
		*aggregationLines = append(*aggregationLines, ColorANSI(opts, optsTypeTitleAggregation.Colors)...)

		// aggregation prefix
		*aggregationLines = append(*aggregationLines, []byte(optsTypeTitleAggregation.Prefix)...)

		// remove first and last line (-2)
		*aggregationLines = append(*aggregationLines, []byte(strconv.Itoa(aggregationCount.Value-2))...)

		// aggregation suffix
		*aggregationLines = append(*aggregationLines, []byte(optsTypeTitleAggregation.Suffix)...)
		*aggregationLines = append(*aggregationLines, []byte("\n")...)

		// reset aggregation color
		*aggregationLines = append(*aggregationLines, ColorReset...)

		// last line
		lastLine := aggregationCount.Lines[len(aggregationCount.Lines)-1]
		*aggregationLines = append(*aggregationLines, bytes.TrimLeft(lastLine, "\n")...)
		return
	}

	// no aggregation output
	for i := 0; i < len(aggregationCount.Lines); i++ {
		*aggregationLines = append(*aggregationLines, aggregationCount.Lines[i]...)
	}
}

func GetOptsTypeTitleAggregationFromAggregationCountType(opts Opts, aggregationCount *AggregationCount) LineTypeTitleAggregation {
	switch aggregationCount.Type {
	// run
	case AggregationTypeRun:
		return opts.Run.Title.Aggregation

	// fail
	case AggregationTypeFail:
		return opts.Fail.Title.Aggregation

	// pass
	case AggregationTypePass:
		return opts.Pass.Title.Aggregation

	// skip
	case AggregationTypeSkip:
		return opts.Skip.Title.Aggregation

	// failed
	case AggregationTypeFailed:
		return opts.Failed.Title.Aggregation

	// ok
	case AggregationTypeOk:
		return opts.Ok.Title.Aggregation

	// error thrown
	case AggregationTypeErrorThrown:
		return opts.ErrorThrown.Title.Aggregation

	default:
		return LineTypeTitleAggregation{}
	}
}

func HandleAggregation(opts Opts, lineType LineType, aggregationCount *AggregationCount, aggregationType AggregationType, aggregationLines *[]byte, formattedLine *[]byte) {
	if aggregationCount.Type != aggregationType {
		optsTypeTitleAggregation := GetOptsTypeTitleAggregationFromAggregationCountType(opts, aggregationCount)
		PrintAggregation(opts, optsTypeTitleAggregation, aggregationCount, aggregationLines)
	}

	if lineType.Title.Aggregation.Activate {
		if aggregationCount.Type != aggregationType {
			aggregationCount.Type = aggregationType
			aggregationCount.Value = 0
			aggregationCount.Lines = [][]byte{*formattedLine}
		} else {
			aggregationCount.Lines = append(aggregationCount.Lines, *formattedLine)
		}

		aggregationCount.Value++
		*formattedLine = nil
		return
	}

	// no aggregation
	if aggregationCount.Type != AggregationTypeNone {
		aggregationCount.Type = AggregationTypeNone
		aggregationCount.Value = 0
		aggregationCount.Lines = [][]byte{}
	}
}

func FormatTestEndLine(line []byte, formattedLine *[]byte, color []byte) {
	if len(color) > 0 {
		*formattedLine = append(*formattedLine, color...)
	}

	*formattedLine = append(*formattedLine, line...)
	*formattedLine = append(*formattedLine, ColorReset...)
	*formattedLine = append(*formattedLine, []byte("\n")...)
}

func FormatTestLine(
	opts Opts,
	line []byte,
	errorBefore *bool,
	stdout io.Writer,
	lineSummary *LineSummary,
	aggregationCount *AggregationCount,
) ([]byte, []byte) {
	var formattedLine []byte
	var aggregationLines []byte

	if len(line) > 0 {
		line = bytes.TrimRight(line, "\n")
		line = bytes.TrimLeft(line, " ")

		var color []byte

		// manage color and style line depending on the content
		// INFO: === RUN
		if bytes.Contains(line, DefaultTitle.Run) {
			if opts.Run.Title.Hide {
				return nil, nil
			}

			if !opts.Run.Summary.Hide {
				lineSummary.Run++
			}

			HandleLineType(opts, &line, opts.Run, DefaultTitle.Run, &color, stdout, errorBefore, false)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Run, aggregationCount, AggregationTypeRun, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: --- FAIL:
		} else if bytes.Contains(line, DefaultTitle.Fail) {
			if opts.Fail.Title.Hide {
				return nil, nil
			}

			if !opts.Fail.Summary.Hide {
				lineSummary.Fail++
			}

			HandleLineType(opts, &line, opts.Fail, DefaultTitle.Fail, &color, stdout, errorBefore, false)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Fail, aggregationCount, AggregationTypeFail, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: --- PASS:
		} else if bytes.Contains(line, DefaultTitle.Pass) {
			if opts.Pass.Title.Hide {
				return nil, nil
			}

			if !opts.Pass.Summary.Hide {
				lineSummary.Pass++
			}

			HandleLineType(opts, &line, opts.Pass, DefaultTitle.Pass, &color, stdout, errorBefore, false)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Pass, aggregationCount, AggregationTypePass, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: --- SKIP:
		} else if bytes.Contains(line, DefaultTitle.Skip) {
			if opts.Skip.Title.Hide {
				return nil, nil
			}

			if !opts.Skip.Summary.Hide {
				lineSummary.Skip++
			}

			HandleLineType(opts, &line, opts.Skip, DefaultTitle.Skip, &color, stdout, errorBefore, false)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Skip, aggregationCount, AggregationTypeSkip, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: FAIL
		} else if bytes.Equal(line, DefaultTitle.Failed) {
			if opts.Failed.Title.Hide {
				return nil, nil
			}

			if !opts.Failed.Summary.Hide {
				lineSummary.Failed++
			}

			HandleLineType(opts, &line, opts.Failed, DefaultTitle.Failed, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Failed, aggregationCount, AggregationTypeFailed, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: ok
		} else if bytes.Equal(line, DefaultTitle.Ok) {
			if opts.Ok.Title.Hide {
				return nil, nil
			}

			if !opts.Ok.Summary.Hide {
				lineSummary.Ok++
			}

			HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, errorBefore, false)
			formattedLine = append(formattedLine, []byte("\n")...)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Ok, aggregationCount, AggregationTypeOk, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines

			// INFO: error thrown
		} else {
			if opts.ErrorThrown.Title.Hide {
				return nil, nil
			}

			if !opts.ErrorThrown.Summary.Hide {
				lineSummary.ErrorThrown++
			}

			HandleLineType(opts, &line, opts.ErrorThrown, DefaultTitle.ErrorThrown, &color, stdout, errorBefore, true)

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.ErrorThrown, aggregationCount, AggregationTypeErrorThrown, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines
		}
	}

	return nil, nil
}

func AddPrintLineSummary(print *[]byte, opts Opts, summary LineTypeSummary, value int) {
	if len(*print) == 0 {
		*print = append(*print, []byte("\n")...)

		// manage header summary
		*print = append(*print, ColorANSI(opts, opts.Summary.Header.Colors)...)
		*print = append(*print, []byte(opts.Summary.Header.Title)...)
		*print = append(*print, ColorReset...)
		*print = append(*print, []byte("\n")...)
	}

	// manage summary content
	*print = append(*print, ColorANSI(opts, summary.Colors)...)
	*print = append(*print, []byte(summary.Prefix)...)
	*print = append(*print, []byte(strconv.Itoa(value))...)
	*print = append(*print, []byte(summary.Suffix)...)

	// manage end of line
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
		print = append(print, ColorANSI(opts, opts.Summary.Footer.Colors)...)
		print = append(print, []byte(opts.Summary.Footer.Title)...)
		print = append(print, ColorReset...)
		print = append(print, []byte("\n")...)
		print = append(print, []byte("\n")...)
		return print
	}

	return nil
}
