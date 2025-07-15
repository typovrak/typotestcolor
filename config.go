package typotestcolor

type LineTypeEnum int

const (
	LineTypeEnumNone LineTypeEnum = iota
	LineTypeEnumRun
	LineTypeEnumFail
	LineTypeEnumPass
	LineTypeEnumSkip
	LineTypeEnumFailed
	LineTypeEnumOk
	LineTypeEnumErrorThrown
)

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

func NewDefaultOpts() Opts {
	return Opts{
		Run: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t=== RUN: ",
				Suffix: "",
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
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Run: ",
				Suffix: "",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER RUN ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundCyan],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER RUN ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundCyan],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
			},
		},
		Fail: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- FAIL: ",
				Suffix: "",
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
				Prefix: "Fail: ",
				Suffix: " ✗",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER FAIL ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundRed],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER FAIL ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundRed],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
			},
		},
		Pass: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- PASS: ",
				Suffix: "",
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
				Prefix: "Pass: ",
				Suffix: " ✓",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER PASS ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundGreen],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER PASS ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundGreen],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
			},
		},
		Skip: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundYellow],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- SKIP: ",
				Suffix: "",
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
				Prefix: "Skip: ",
				Suffix: " ~",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER SKIP ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundYellow],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER SKIP ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundYellow],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
			},
		},
		Failed: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundRed],
				},
				Prefix: "\nFAIL",
				Suffix: "",
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
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Failed: ",
				Suffix: " ✗",
				Hide:   true,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER FAILED ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER FAILED ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide: false,
				},
			},
		},
		Ok: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundGreen],
				},
				Prefix: "\nPASS",
				Suffix: "",
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
					Style:      ColorANSISTyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Ok: ",
				Suffix: " ✓",
				Hide:   true,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER OK ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER OK ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide: false,
				},
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
				Suffix: "",
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
				Prefix: "Error thrown: ",
				Suffix: " !",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "--- HEADER ERROR THROWN ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundWhite],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "--- FOOTER ERROR THROWN ---",
					Colors: ANSIConfig{
						Style:      ColorANSISTyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundWhite],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Hide: false,
				},
			},
		},
		Debug: false,
		Summary: SummaryConfig{
			Header: SummaryConfigTitle{
				Title: "\n-----> TESTS SUMMARY <-----",
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundPurple],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
			},
			Footer: SummaryConfigTitle{
				Title: "-----> Made with <3 by Typovrak <-----\n",
				Colors: ANSIConfig{
					Style:      ColorANSISTyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundPurple],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
			},
			AlignResults: true,
		},
	}
}
