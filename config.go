package typotestcolor

type Opts struct {
	Run         LineType
	Fail        LineType
	Pass        LineType
	Skip        LineType
	Failed      LineType
	Ok          LineType
	ErrorThrown LineType
}

type LineType struct {
	Colors ANSIConfig
	Title  string
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
		},
		Fail: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundRed],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- FAIL:",
		},
		Pass: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundGreen],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- PASS:",
		},
		Skip: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundYellow],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "\t--- SKIP:",
		},
		Failed: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundRed],
			},
			Title: "FAIL",
		},
		Ok: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundGreen],
			},
			Title: "PASS",
		},

		ErrorThrown: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundWhite],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "",
		},
	}
}
