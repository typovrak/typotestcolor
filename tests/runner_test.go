package tests

import (
	"testing"
)

func TestRunTestColor(t *testing.T) {
	//	testOutput1 := `=== RUN   TestColorANSI
	//=== RUN   TestColorANSI/run_default_color
	//=== RUN   TestColorANSI/fail_default_color
	//=== RUN   TestColorANSI/pass_default_color
	//=== RUN   TestColorANSI/skip_default_color
	//=== RUN   TestColorANSI/failed_default_color
	//=== RUN   TestColorANSI/ok_default_color
	//=== RUN   TestColorANSI/error_thrown_default_color
	//--- PASS: TestColorANSI (0.00s)
	//    --- PASS: TestColorANSI/run_default_color (0.00s)
	//    --- PASS: TestColorANSI/fail_default_color (0.00s)
	//    --- PASS: TestColorANSI/pass_default_color (0.00s)
	//    --- PASS: TestColorANSI/skip_default_color (0.00s)
	//    --- PASS: TestColorANSI/failed_default_color (0.00s)
	//    --- PASS: TestColorANSI/ok_default_color (0.00s)
	//    --- PASS: TestColorANSI/error_thrown_default_color (0.00s)`

	t.Run("default", func(t *testing.T) {
		// r, w, _ := os.Pipe()
		// os.Stdout = w

		// exitCode := typotestcolor.RunTestColor(nil, typotestcolor.NewDefaultOpts())

		// w.Close()
		// out, _ := io.ReadAll(r)

		// os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

		//if exitCode != 0 {
		//	t.Errorf("[exitCode] expected %d, got %d", 0, exitCode)
		//}
	})
}
