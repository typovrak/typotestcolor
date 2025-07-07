package main

import (
	"bytes"
	"os"
)

func addLineFeedBetweenErrorThrown(w *os.File, errorBefore *bool, isError bool) {
	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

func handleLineType(
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w *os.File,
	errorBefore *bool,
	isError bool,
) {
	*color = ColorANSI(lineType.Colors)
	addLineFeedBetweenErrorThrown(w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
}
