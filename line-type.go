package typotestcolor

import (
	"bytes"
	"io"
)

type LineType struct {
	Title   LineTypeTitle
	Summary LineTypeSummary
	Section LineTypeSection
}

type LineTypeTitle struct {
	Colors      ANSIConfig
	Prefix      string
	Suffix      string
	Hide        bool
	Aggregation LineTypeTitleAggregation
}

type LineTypeSection struct {
	Header LineTypeSectionTitle
	Footer LineTypeSectionTitle
}

type LineTypeSectionTitle struct {
	Title  string
	Colors ANSIConfig
	Hide   bool
}

func AddLineFeedBetweenErrorThrown(opts Opts, w io.Writer, lineTypeBefore *LineTypeEnum, isError bool) {
	Debug(opts, "AddLineFeedBetweenErrorThrown")

	if (!isError && *lineTypeBefore == LineTypeEnumErrorThrown) || (isError && *lineTypeBefore != LineTypeEnumErrorThrown) {
		w.Write([]byte("\n"))
	}
}

func HandleLineType(
	opts Opts,
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w io.Writer,
	lineTypeBefore *LineTypeEnum,
	isError bool,
) {
	Debug(opts, "HandleLineType")

	*color = ColorANSI(opts, lineType.Title.Colors)
	AddLineFeedBetweenErrorThrown(opts, w, lineTypeBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title.Prefix), 1)
	*line = append(*line, []byte(lineType.Title.Suffix)...)
}

func FormatTestEndLine(line []byte, formattedLine *[]byte, color []byte) {
	if len(color) > 0 {
		*formattedLine = append(*formattedLine, color...)
	}

	*formattedLine = append(*formattedLine, line...)
	*formattedLine = append(*formattedLine, ColorReset...)
	*formattedLine = append(*formattedLine, '\n')
}

func HandleSectionHeader(opts Opts, lineType LineType, formattedLine *[]byte, lineTypeBefore *LineTypeEnum, lineTypeCurrent LineTypeEnum) {
	if lineType.Section.Header.Hide ||
		*lineTypeBefore == lineTypeCurrent {
		return
	}

	*formattedLine = append(*formattedLine, ColorANSI(opts, lineType.Section.Header.Colors)...)
	*formattedLine = append(*formattedLine, []byte(lineType.Section.Header.Title)...)
	*formattedLine = append(*formattedLine, ColorReset...)
	*formattedLine = append(*formattedLine, '\n')
}

func HandleSectionFooter(opts Opts, formattedLine *[]byte, lineTypeBefore *LineTypeEnum, lineTypeCurrent LineTypeEnum) {
	var previousLineType LineType

	switch *lineTypeBefore {
	// run
	case LineTypeEnumRun:
		previousLineType = opts.Run

		// fail
	case LineTypeEnumFail:
		previousLineType = opts.Fail

		// pass
	case LineTypeEnumPass:
		previousLineType = opts.Pass

		// skip
	case LineTypeEnumSkip:
		previousLineType = opts.Skip

		// failed
	case LineTypeEnumFailed:
		previousLineType = opts.Failed

		// ok
	case LineTypeEnumOk:
		previousLineType = opts.Ok

		// error thrown
	case LineTypeEnumErrorThrown:
		previousLineType = opts.ErrorThrown

	}

	if previousLineType.Section.Footer.Hide ||
		*lineTypeBefore == LineTypeEnumNone ||
		*lineTypeBefore == lineTypeCurrent {
		return
	}

	*formattedLine = append(*formattedLine, ColorANSI(opts, previousLineType.Section.Footer.Colors)...)
	*formattedLine = append(*formattedLine, []byte(previousLineType.Section.Footer.Title)...)
	*formattedLine = append(*formattedLine, ColorReset...)
	*formattedLine = append(*formattedLine, '\n')
}

func FormatTestLine(
	opts Opts,
	line []byte,
	lineTypeBefore *LineTypeEnum,
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

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumRun)
			HandleSectionHeader(opts, opts.Run, &formattedLine, lineTypeBefore, LineTypeEnumRun)

			HandleLineType(opts, &line, opts.Run, DefaultTitle.Run, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Run, aggregationCount, LineTypeEnumRun, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumRun

			// INFO: --- FAIL:
		} else if bytes.Contains(line, DefaultTitle.Fail) {
			if opts.Fail.Title.Hide {
				return nil, nil
			}

			if !opts.Fail.Summary.Hide {
				lineSummary.Fail++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumFail)
			HandleSectionHeader(opts, opts.Fail, &formattedLine, lineTypeBefore, LineTypeEnumFail)

			HandleLineType(opts, &line, opts.Fail, DefaultTitle.Fail, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Fail, aggregationCount, LineTypeEnumFail, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumFail

			// INFO: --- PASS:
		} else if bytes.Contains(line, DefaultTitle.Pass) {
			if opts.Pass.Title.Hide {
				return nil, nil
			}

			if !opts.Pass.Summary.Hide {
				lineSummary.Pass++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumPass)
			HandleSectionHeader(opts, opts.Pass, &formattedLine, lineTypeBefore, LineTypeEnumPass)

			HandleLineType(opts, &line, opts.Pass, DefaultTitle.Pass, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Pass, aggregationCount, LineTypeEnumPass, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumPass

			// INFO: --- SKIP:
		} else if bytes.Contains(line, DefaultTitle.Skip) {
			if opts.Skip.Title.Hide {
				return nil, nil
			}

			if !opts.Skip.Summary.Hide {
				lineSummary.Skip++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumSkip)
			HandleSectionHeader(opts, opts.Skip, &formattedLine, lineTypeBefore, LineTypeEnumSkip)

			HandleLineType(opts, &line, opts.Skip, DefaultTitle.Skip, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Skip, aggregationCount, LineTypeEnumSkip, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumSkip

			// INFO: FAIL
		} else if bytes.Equal(line, DefaultTitle.Failed) {
			if opts.Failed.Title.Hide {
				return nil, nil
			}

			if !opts.Failed.Summary.Hide {
				lineSummary.Failed++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumFailed)
			HandleSectionHeader(opts, opts.Failed, &formattedLine, lineTypeBefore, LineTypeEnumFailed)

			HandleLineType(opts, &line, opts.Failed, DefaultTitle.Failed, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Failed, aggregationCount, LineTypeEnumFailed, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumFailed

			// INFO: ok
		} else if bytes.Equal(line, DefaultTitle.Ok) {
			if opts.Ok.Title.Hide {
				return nil, nil
			}

			if !opts.Ok.Summary.Hide {
				lineSummary.Ok++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumOk)
			HandleSectionHeader(opts, opts.Ok, &formattedLine, lineTypeBefore, LineTypeEnumOk)

			HandleLineType(opts, &line, opts.Ok, DefaultTitle.Ok, &color, stdout, lineTypeBefore, false)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.Ok, aggregationCount, LineTypeEnumOk, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumOk

			// INFO: error thrown
		} else {
			if opts.ErrorThrown.Title.Hide {
				return nil, nil
			}

			if !opts.ErrorThrown.Summary.Hide {
				lineSummary.ErrorThrown++
			}

			HandleSectionFooter(opts, &formattedLine, lineTypeBefore, LineTypeEnumErrorThrown)
			HandleSectionHeader(opts, opts.ErrorThrown, &formattedLine, lineTypeBefore, LineTypeEnumErrorThrown)

			HandleLineType(opts, &line, opts.ErrorThrown, DefaultTitle.ErrorThrown, &color, stdout, lineTypeBefore, true)
			FormatTestEndLine(line, &formattedLine, color)
			HandleAggregation(opts, opts.ErrorThrown, aggregationCount, LineTypeEnumErrorThrown, &aggregationLines, &formattedLine)

			*lineTypeBefore = LineTypeEnumErrorThrown
		}
	}

	return formattedLine, aggregationLines
}
