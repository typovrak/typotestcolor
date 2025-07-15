package typotestcolor

import "strconv"

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
}

type SummaryConfigTitle = struct {
	Title  string
	Colors ANSIConfig
}

func AddPrintLineSummary(print *[]byte, opts Opts, summary LineTypeSummary, value int) {
	if len(*print) == 0 {
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

type SummaryData = struct {
	Type   int
	Prefix string
	Value  int
	Suffix string
}

func PrintLineSummary(opts Opts, lineSummary LineSummary) []byte {
	//	var data SummaryData
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
		return print
	}

	return nil
}
