package typotestcolor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func DiffPrintColor(color ANSIForeground, highlight bool, opts Opts) []byte {
	var (
		print          []byte
		ansiConfig     ANSIConfig
		ansiForeground ANSIForeground
		ansiBackground ANSIBackground
	)

	switch color {
	case ANSIForegroundGreen:
		ansiForeground = ANSIForegroundGreen
		ansiBackground = ANSIBackgroundGreen

	case ANSIForegroundRed:
		ansiForeground = ANSIForegroundRed
		ansiBackground = ANSIBackgroundRed

	default:
		ansiForeground = ANSIForegroundBlue
		ansiBackground = ANSIBackgroundBlue
	}

	if highlight {
		ansiConfig = ANSIConfig{
			Style:      ColorANSIStyle[ANSIStyleNormal],
			Foreground: ColorANSIForeground[ANSIForegroundBlack],
			Background: ColorANSIBackground[ansiBackground],
		}
	} else {
		ansiConfig = ANSIConfig{
			Style:      ColorANSIStyle[ANSIStyleNormal],
			Foreground: ColorANSIForeground[ansiForeground],
			Background: ColorANSIBackground[ANSIBackgroundNone],
		}

		// INFO: need to reset for ANSIBackgroundNone to works
		print = append(print, ColorReset...)
	}

	print = append(print, ColorANSI(opts, ansiConfig)...)
	return print
}

func ToBytes(value any) ([]byte, error) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() == reflect.Func && reflectValue.Type().NumIn() == 0 {
		results := reflectValue.Call(nil)

		if len(results) > 0 {
			return ToBytes(results[0].Interface())
		}
	}

	switch valueTyped := value.(type) {

	case []byte:
		return valueTyped, nil

	case string:
		return []byte(valueTyped), nil

	case int, int8, int16, int32, int64:
		return fmt.Appendf(nil, "%d", valueTyped), nil

	case uint, uint8, uint16, uint32, uint64:
		return fmt.Appendf(nil, "%d", valueTyped), nil

	case float32, float64:
		return fmt.Appendf(nil, "%g", valueTyped), nil

	case bool:
		return []byte(strconv.FormatBool(valueTyped)), nil

	case nil:
		return []byte("nil"), nil

	case fmt.Stringer:
		return []byte(valueTyped.String()), nil

	case func() []byte:
		return valueTyped(), nil

	case func() string:
		return []byte(valueTyped()), nil

	// INFO: cannot aggregate func() type
	case func() int:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() int8:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() int16:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() int32:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() int64:
		return fmt.Appendf(nil, "%d", valueTyped()), nil

	case func() uint:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() uint8:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() uint16:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() uint32:
		return fmt.Appendf(nil, "%d", valueTyped()), nil
	case func() uint64:
		return fmt.Appendf(nil, "%d", valueTyped()), nil

	case func() float32:
		return fmt.Appendf(nil, "%g", valueTyped()), nil
	case func() float64:
		return fmt.Appendf(nil, "%g", valueTyped()), nil

	case func() bool:
		return []byte(strconv.FormatBool(valueTyped())), nil

	case func():
		return []byte("nil"), nil

	case func() fmt.Stringer:
		return []byte(valueTyped().String()), nil

	default:
		fmt.Println("in default")
		// fallback JSON (struct, map, slice, etc.)
		json, err := json.Marshal(valueTyped)
		if err != nil {
			return nil, fmt.Errorf("cannot convert %T to []byte: %w", valueTyped, err)
		}

		return json, nil
	}
}

type TestDiffOpts = struct {
	opts         Opts
	printToASCII bool
}

func TestDiffNewDefaultOpts() TestDiffOpts {
	return TestDiffOpts{
		opts:         NewDefaultOpts(),
		printToASCII: false,
	}
}

func TestDiffDefault(t *testing.T, expected any, got any) {
	TestDiff(t, expected, got, TestDiffNewDefaultOpts())
}

func TestDiff(t *testing.T, expected any, got any, opts TestDiffOpts) {
	expectedBytes, err := ToBytes(expected)
	if err != nil {
		t.Error(err)
		return
	}

	gotBytes, err := ToBytes(got)
	if err != nil {
		t.Error(err)
		return
	}

	if bytes.Equal(expectedBytes, gotBytes) {
		return
	}

	var print strings.Builder

	expectedLen := len(expectedBytes)
	gotLen := len(gotBytes)

	expectedPrefix := fmt.Sprintf("expected (length: %d): ", expectedLen)
	gotPrefix := fmt.Sprintf("got (length: %d): ", gotLen)

	expectedPrefixLen := len(expectedPrefix)
	gotPrefixLen := len(gotPrefix)

	if expectedPrefixLen > gotPrefixLen {
		for i := 0; i < expectedPrefixLen-gotPrefixLen; i++ {
			gotPrefix += " "
		}
	} else if gotPrefixLen > expectedPrefixLen {
		for i := 0; i < gotPrefixLen-expectedPrefixLen; i++ {
			expectedPrefix += " "
		}
	}

	// start header print
	print.WriteByte('\n')

	// expected part
	print.WriteString(expectedPrefix)
	print.Write(DiffPrintColor(ANSIForegroundGreen, false, opts.opts))
	print.Write(expectedBytes)
	print.WriteByte('\n')

	// got part
	print.WriteString(gotPrefix)
	print.Write(DiffPrintColor(ANSIForegroundRed, false, opts.opts))

	errorBefore := false

	// build got print
	for i := 0; i < min(expectedLen, gotLen); i++ {
		// same
		if expectedBytes[i] == gotBytes[i] {
			// remove error highlight
			if errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, false, opts.opts))
				errorBefore = false
			}

			// diff
		} else {
			// add error highlight
			if !errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, true, opts.opts))
				errorBefore = true
			}
		}

		print.WriteByte(gotBytes[i])
	}

	// add got trailing bytes
	if expectedLen < gotLen {
		if !errorBefore {
			print.Write(DiffPrintColor(ANSIForegroundRed, true, opts.opts))
			errorBefore = true
		}

		print.Write(gotBytes[expectedLen:])
	}

	// reset ANSI styles
	print.Write(ColorReset)
	print.WriteByte('\n')

	if opts.printToASCII {
		t.Error(strconv.QuoteToASCII(print.String()))
	} else {
		t.Error(print.String())
	}
}
