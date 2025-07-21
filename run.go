package typotestcolor

import "strconv"

func RunTitle(index *int, message string) string {
	title := strconv.Itoa(*index) + "_" + message
	*index++
	return title
}
