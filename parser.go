package typotestcolor

import (
	"bytes"
	"io"
)

func AddLineFeedBetweenErrorThrown(w io.Writer, errorBefore *bool, isError bool) {
	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

func HandleLineType(
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w io.Writer,
	errorBefore *bool,
	isError bool,
) {
	*color = ColorANSI(lineType.Colors)
	AddLineFeedBetweenErrorThrown(w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
}
