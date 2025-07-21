package typotestcolor

import (
	"strconv"
)

type ANSIStyle int

const (
	ANSIStyleReset ANSIStyle = iota
	ANSIStyleBold
	ANSIStyleDim
	ANSIStyleUnderline
	ANSIStyleInverse
	ANSIStyleHidden
	ANSIStyleNormal
)

var ColorANSIStyle = map[ANSIStyle]int{
	ANSIStyleReset:     0,
	ANSIStyleBold:      1,
	ANSIStyleDim:       2,
	ANSIStyleUnderline: 4,
	ANSIStyleInverse:   7,
	ANSIStyleHidden:    8,
	ANSIStyleNormal:    22,
}

type ANSIForeground int

const (
	ANSIForegroundNone ANSIForeground = iota
	ANSIForegroundBlack
	ANSIForegroundRed
	ANSIForegroundGreen
	ANSIForegroundYellow
	ANSIForegroundBlue
	ANSIForegroundPurple
	ANSIForegroundCyan
	ANSIForegroundWhite
)

var ColorANSIForeground = map[ANSIForeground]int{
	ANSIForegroundNone:   0,
	ANSIForegroundBlack:  30,
	ANSIForegroundRed:    31,
	ANSIForegroundGreen:  32,
	ANSIForegroundYellow: 33,
	ANSIForegroundBlue:   34,
	ANSIForegroundPurple: 35,
	ANSIForegroundCyan:   36,
	ANSIForegroundWhite:  37,
}

type ANSIBackground int

const (
	ANSIBackgroundNone ANSIBackground = iota
	ANSIBackgroundBlack
	ANSIBackgroundRed
	ANSIBackgroundGreen
	ANSIBackgroundYellow
	ANSIBackgroundBlue
	ANSIBackgroundPurple
	ANSIBackgroundCyan
	ANSIBackgroundWhite
)

var ColorANSIBackground = map[ANSIBackground]int{
	ANSIBackgroundNone:   0,
	ANSIBackgroundBlack:  40,
	ANSIBackgroundRed:    41,
	ANSIBackgroundGreen:  42,
	ANSIBackgroundYellow: 43,
	ANSIBackgroundBlue:   44,
	ANSIBackgroundPurple: 45,
	ANSIBackgroundCyan:   46,
	ANSIBackgroundWhite:  47,
}

// INFO: always use ColorReset before adding a line feed (\n)
var ColorReset = []byte("\033[0m")

type ANSIConfig struct {
	Style      int
	Foreground int
	Background int
}

func ColorANSI(config ANSIConfig) []byte {
	color := []byte("\033[")

	color = append(color, []byte(strconv.Itoa(config.Style))...)
	color = append(color, ';')
	color = append(color, []byte(strconv.Itoa(config.Foreground))...)

	if config.Background != ColorANSIBackground[ANSIBackgroundNone] {
		color = append(color, ';')
		color = append(color, []byte(strconv.Itoa(config.Background))...)
	}

	color = append(color, 'm')

	return color
}
