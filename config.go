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
	Colors ANSIConfig
	Title  string
	Hide   bool
}

func NewDefaultOpts() Opts {
	return Opts{
		Run: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundCyan],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t=== RUN:",
			Hide:  false,
		},
		Fail: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundRed],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- FAIL:",
			Hide:  false,
		},
		Pass: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundGreen],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- PASS:",
			Hide:  false,
		},
		Skip: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundYellow],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- SKIP:",
			Hide:  false,
		},
		Failed: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundRed],
			},
			Title: "FAIL",
			Hide:  false,
		},
		Ok: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundGreen],
			},
			Title: "PASS",
			Hide:  false,
		},

		ErrorThrown: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundWhite],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "",
			Hide:  false,
		},
		Debug: false,
	}
}
