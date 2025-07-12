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
}

type LineType struct {
	Colors           ANSIConfig
	Title            string
	Hide             bool
	AggregationTitle string
	AggregationHide  bool
}

type LineAggregation = struct {
	Run         int
	Fail        int
	Pass        int
	Skip        int
	Failed      int
	Ok          int
	ErrorThrown int
}

func NewDefaultOpts() Opts {
	return Opts{
		Run: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundCyan],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title:            "\t=== RUN:",
			Hide:             false,
			AggregationTitle: "RUN:",
			AggregationHide:  false,
		},
		Fail: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundRed],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title:            "\t--- FAIL:",
			Hide:             false,
			AggregationTitle: "FAIL:",
			AggregationHide:  false,
		},
		Pass: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundGreen],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title:            "\t--- PASS:",
			Hide:             false,
			AggregationTitle: "PASS:",
			AggregationHide:  false,
		},
		Skip: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundYellow],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title:            "\t--- SKIP:",
			Hide:             false,
			AggregationTitle: "SKIP:",
			AggregationHide:  false,
		},
		Failed: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundRed],
			},
			Title:            "FAIL",
			Hide:             false,
			AggregationTitle: "FAILED:",
			AggregationHide:  false,
		},
		Ok: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundGreen],
			},
			Title:            "PASS",
			Hide:             false,
			AggregationTitle: "OK:",
			AggregationHide:  false,
		},

		ErrorThrown: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundWhite],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title:            "",
			Hide:             false,
			AggregationTitle: "ERROR_THROWN:",
			AggregationHide:  false,
		},
		Debug: false,
	}
}
