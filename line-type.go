package typotestcolor

import (
	"bytes"
	"io"
)

type LineType struct {
	Title   LineTypeTitle
	Summary LineTypeSummary
}

type LineTypeTitle struct {
	Colors      ANSIConfig
	Prefix      string
	Suffix      string
	Hide        bool
	Aggregation LineTypeTitleAggregation
}

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
	*line = append(*line, []byte(lineType.Title.Suffix)...)
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

			HandleAggregation(opts, opts.Run, aggregationCount, LineTypeEnumRun, &aggregationLines, &formattedLine)
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

			HandleAggregation(opts, opts.Fail, aggregationCount, LineTypeEnumFail, &aggregationLines, &formattedLine)
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

			HandleAggregation(opts, opts.Pass, aggregationCount, LineTypeEnumPass, &aggregationLines, &formattedLine)
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

			HandleAggregation(opts, opts.Skip, aggregationCount, LineTypeEnumSkip, &aggregationLines, &formattedLine)
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

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Failed, aggregationCount, LineTypeEnumFailed, &aggregationLines, &formattedLine)
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

			FormatTestEndLine(line, &formattedLine, color)

			HandleAggregation(opts, opts.Ok, aggregationCount, LineTypeEnumOk, &aggregationLines, &formattedLine)
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

			HandleAggregation(opts, opts.ErrorThrown, aggregationCount, LineTypeEnumErrorThrown, &aggregationLines, &formattedLine)
			return formattedLine, aggregationLines
		}
	}

	return nil, nil
}
