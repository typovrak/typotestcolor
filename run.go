package typotestcolor

import "strconv"

// TODO: cette fonction ne résoud pas mon problème de grep...
func RunTitle(index *int, message string) string {
	title := strconv.Itoa(*index) + "_" + message
	*index++
	return title
}
