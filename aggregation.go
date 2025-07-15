package typotestcolor

import (
	"bytes"
	"strconv"
)

type AggregationType int

const (
	AggregationTypeNone AggregationType = iota
	AggregationTypeRun
	AggregationTypeFail
	AggregationTypePass
	AggregationTypeSkip
	AggregationTypeFailed
	AggregationTypeOk
	AggregationTypeErrorThrown
)

type AggregationCount = struct {
	Type  AggregationType
	Value int
	Lines [][]byte
}

type LineTypeTitleAggregation = struct {
	Activate  bool
	Colors    ANSIConfig
	Prefix    string
	Suffix    string
	Threshold int
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
