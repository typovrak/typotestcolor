package typotestcolor

import "fmt"

func Debug(opts Opts, funcName string) {
	if opts.Debug {
		fmt.Println("DEBUG:", funcName)
	}
}
