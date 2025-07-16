package typotestcolor

type Opts struct {
	Run                  LineType
	Fail                 LineType
	Pass                 LineType
	Skip                 LineType
	Failed               LineType
	Ok                   LineType
	ErrorThrown          LineType
	Debug                bool
	Summary              SummaryConfig
	SeparateEverySection SeparateEverySectionStruct
}

type SeparateEverySectionStruct = struct {
	Colors ANSIConfig
	Title  string
	Hide   bool
}

func NewDefaultOpts() Opts {
	return Opts{
		Run: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t=== RUN: ",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundCyan],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundCyan],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Run: ",
				Suffix: "",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t RUN section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundCyan],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "\t RUN section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundCyan],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
			},
		},
		Fail: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- FAIL: ",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					// INFO: important to see every fail case
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundRed],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Fail: ",
				Suffix: " ✗",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t FAIL section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "\t FAIL section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
			},
		},
		Pass: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- PASS: ",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundGreen],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Pass: ",
				Suffix: " ✓",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t PASS section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "\t PASS section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
			},
		},
		Skip: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundYellow],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "\t--- SKIP: ",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: true,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundYellow],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 4,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundYellow],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Skip: ",
				Suffix: " ~",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t SKIP section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundYellow],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
				Footer: LineTypeSectionTitle{
					Title: "\t SKIP section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundYellow],
					},
					Hide:             false,
					AddEmptyLineFeed: false,
				},
			},
		},
		Failed: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundRed],
				},
				Prefix: "\n\n Tests result: FAIL\n",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundRed],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Failed: ",
				Suffix: " ✗",
				Hide:   true,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t FAILED section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide:             true,
					AddEmptyLineFeed: false,
				},
				// INFO: can't be displayed (and useless)
				Footer: LineTypeSectionTitle{
					Title: "\t FAILED section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundRed],
					},
					Hide:             true,
					AddEmptyLineFeed: true,
				},
			},
		},
		Ok: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundGreen],
				},
				Prefix: "\n\n Tests result: PASS\n",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundGreen],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Ok: ",
				Suffix: " ✓",
				Hide:   true,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t OK section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide:             true,
					AddEmptyLineFeed: false,
				},
				// INFO: can't be displayed (and useless)
				Footer: LineTypeSectionTitle{
					Title: "\t OK section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundGreen],
					},
					Hide:             true,
					AddEmptyLineFeed: true,
				},
			},
		},
		ErrorThrown: LineType{
			Title: LineTypeTitle{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundWhite],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "",
				Suffix: "",
				Hide:   false,
				Aggregation: LineTypeTitleAggregation{
					Activate: false,
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleNormal],
						Foreground: ColorANSIForeground[ANSIForegroundWhite],
						Background: ColorANSIBackground[ANSIBackgroundNone],
					},
					Prefix:    "\t[",
					Suffix:    "]",
					Threshold: 0,
				},
			},
			Summary: LineTypeSummary{
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleNormal],
					Foreground: ColorANSIForeground[ANSIForegroundWhite],
					Background: ColorANSIBackground[ANSIBackgroundNone],
				},
				Prefix: "Error thrown: ",
				Suffix: " !",
				Hide:   false,
			},
			Section: LineTypeSection{
				Header: LineTypeSectionTitle{
					Title: "\t ERROR THROWN section start ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundWhite],
					},
					Hide:             false,
					AddEmptyLineFeed: true,
				},
				Footer: LineTypeSectionTitle{
					Title: "\t ERROR THROWN section end ",
					Colors: ANSIConfig{
						Style:      ColorANSIStyle[ANSIStyleBold],
						Foreground: ColorANSIForeground[ANSIForegroundBlack],
						Background: ColorANSIBackground[ANSIBackgroundWhite],
					},
					Hide:             false,
					AddEmptyLineFeed: true,
				},
			},
		},
		Debug: false,
		Summary: SummaryConfig{
			Header: SummaryConfigTitle{
				Title: " TESTS SUMMARY ",
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundPurple],
				},
				Hide: false,
			},
			Footer: SummaryConfigTitle{
				Title: "\n Made with <3 by Typovrak ",
				Colors: ANSIConfig{
					Style:      ColorANSIStyle[ANSIStyleBold],
					Foreground: ColorANSIForeground[ANSIForegroundBlack],
					Background: ColorANSIBackground[ANSIBackgroundBlue],
				},
				Hide: false,
			},
			AlignResults: true,
			Hide:         false,
		},
		SeparateEverySection: SeparateEverySectionStruct{
			Colors: ANSIConfig{
				Style:      ColorANSIStyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundNone],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\n",
			Hide:  true,
		},
	}
}
