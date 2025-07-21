package typotestcolor

import (
	"strconv"
)

type LineTypeSummary struct {
	Colors ANSIConfig
	Prefix string
	Suffix string
	Hide   bool
}

type LineSummary = struct {
	Run         int
	Fail        int
	Pass        int
	Skip        int
	Failed      int
	Ok          int
	ErrorThrown int
}

type SummaryConfig = struct {
	Header       SummaryConfigTitle
	Footer       SummaryConfigTitle
	AlignResults bool
	Hide         bool
}

type SummaryConfigTitle = struct {
	Title  string
	Colors ANSIConfig
	Hide   bool
}

func AddPrintLineSummaryHeader(print *[]byte, opts Opts) {
	if opts.Summary.Header.Hide {
		return
	}

	// manage header summary
	*print = append(*print, ColorANSI(opts.Summary.Header.Colors)...)
	*print = append(*print, []byte(opts.Summary.Header.Title)...)
	*print = append(*print, ColorReset...)
	*print = append(*print, '\n')
}

func AddPrintLineSummaryFooter(print *[]byte, opts Opts) {
	if opts.Summary.Footer.Hide {
		return
	}

	*print = append(*print, ColorANSI(opts.Summary.Footer.Colors)...)
	*print = append(*print, []byte(opts.Summary.Footer.Title)...)
	*print = append(*print, ColorReset...)
	*print = append(*print, '\n')
}

func AddPrintLineSummary(print *[]byte, opts Opts, data SummaryData, maxPrefixLen int) {
	// manage summary content
	*print = append(*print, ColorANSI(data.Colors)...)
	*print = append(*print, []byte(data.Prefix)...)

	// align results
	if opts.Summary.AlignResults && len(data.Prefix) < maxPrefixLen {
		for i := 0; i < maxPrefixLen-len(data.Prefix); i++ {
			*print = append(*print, ' ')
		}
	}

	*print = append(*print, []byte(strconv.Itoa(data.Value))...)
	*print = append(*print, []byte(data.Suffix)...)

	// manage end of line
	*print = append(*print, ColorReset...)
	*print = append(*print, '\n')
}

type SummaryData = struct {
	Type   LineTypeEnum
	Colors ANSIConfig
	Prefix string
	Value  int
	Suffix string
}

func PrintLineSummary(opts Opts, lineSummary LineSummary) []byte {
	if opts.Summary.Hide {
		return nil
	}

	var data []SummaryData
	maxPrefixLen := 0

	if !opts.Run.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumRun, opts.Run.Summary, lineSummary.Run)
	}

	if !opts.Fail.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumFail, opts.Fail.Summary, lineSummary.Fail)
	}

	if !opts.Pass.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumPass, opts.Pass.Summary, lineSummary.Pass)
	}

	if !opts.Skip.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumSkip, opts.Skip.Summary, lineSummary.Skip)
	}

	if !opts.Failed.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumFailed, opts.Failed.Summary, lineSummary.Failed)
	}

	if !opts.Ok.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumOk, opts.Ok.Summary, lineSummary.Ok)
	}

	if !opts.ErrorThrown.Summary.Hide {
		AddSummaryData(&data, &maxPrefixLen, LineTypeEnumErrorThrown, opts.ErrorThrown.Summary, lineSummary.ErrorThrown)
	}

	if len(data) == 0 {
		return nil
	}

	// build summary
	var print []byte

	// header
	AddPrintLineSummaryHeader(&print, opts)

	// content
	for i := 0; i < len(data); i++ {
		AddPrintLineSummary(&print, opts, data[i], maxPrefixLen)
	}

	// footer
	AddPrintLineSummaryFooter(&print, opts)

	print = append(print, '\n')
	return print
}

func AddSummaryData(data *[]SummaryData, maxPrefixLen *int, summaryType LineTypeEnum, summary LineTypeSummary, value int) {
	var currentData SummaryData
	currentData.Type = summaryType
	currentData.Colors = summary.Colors
	currentData.Prefix = summary.Prefix
	currentData.Value = value
	currentData.Suffix = summary.Suffix

	if *maxPrefixLen < len(currentData.Prefix) {
		*maxPrefixLen = len(currentData.Prefix)
	}

	*data = append(*data, currentData)
}
