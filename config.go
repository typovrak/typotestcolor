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
		},
	}
}
