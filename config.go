package typotestcolor

type Opts struct {
	Run         LineType
	Fail        LineType
	Pass        LineType
	Skip        LineType
	Failed      LineType
	Ok          LineType
	ErrorThrown LineType
	Debug       bool
	Summary     SummaryConfig
}

type LineType struct {
	Title   LineTypeTitle
	Summary LineTypeSummary
}

type LineTypeTitle struct {
	Colors      ANSIConfig
	Prefix      string
	Hide        bool
	Aggregation LineTypeTitleAggregation
}

type LineTypeSummary struct {
	Colors ANSIConfig
	Prefix string
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

type SummaryConfig = struct {
	Header string
	Footer string
}

func NewDefaultOpts() Opts {
	return Opts{
		Run: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t=== RUN:",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundCyan],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "RUN:",
				Hide:   false,
			},
		},
		Fail: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- FAIL:",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundRed],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "FAIL:",
				Hide:   false,
			},
		},
		Pass: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- PASS:",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundGreen],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "PASS:",
				Hide:   false,
			},
		},
		Skip: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundYellow],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- SKIP:",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundYellow],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundYellow],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "SKIP:",
				Hide:   false,
			},
		},
		Failed: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundRed],
				},
				Prefix: "FAIL",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundRed],
				},
				Prefix: "FAILED:",
				Hide:   false,
			},
		},
		Ok: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundGreen],
				},
				Prefix: "PASS",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundGreen],
				},
				Prefix: "OK:",
				Hide:   false,
			},
		},

		ErrorThrown: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundWhite],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundWhite],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t-> [",
					Suffix:    "] <-",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundWhite],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "ERROR_THROWN:",
				Hide:   false,
			},
		},
		Debug: false,
		Summary: SummaryConfig{
			Header: "----- SUMMARY -----",
			Footer: "----- Made with <3 by Typovrak -----",
		},
	}
}
