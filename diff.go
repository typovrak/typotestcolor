package typotestcolor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func DiffPrintColor(color ANSIForeground, highlight bool) []byte {
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

	print = append(print, ColorANSI(ansiConfig)...)
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

// TODO: mettre printToASCII dans la configuration globale pour simplifier le tout.

func TestDiffDefault(expected any, got any) error {
	return TestDiff(expected, got, TestDiffNewDefaultOpts())
}

// TODO: colorier la valeur length et mettre en highlight la différence

func TestDiff(expected any, got any, opts TestDiffOpts) error {
	expectedBytes, err := ToBytes(expected)
	if err != nil {
		return err
	}

	gotBytes, err := ToBytes(got)
	if err != nil {
		return err
	}

	// everything is fine
	if bytes.Equal(expectedBytes, gotBytes) {
		return nil
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
	print.Write(DiffPrintColor(ANSIForegroundGreen, false))
	print.Write(expectedBytes)
	print.WriteByte('\n')

	// got part
	print.WriteString(gotPrefix)
	print.Write(DiffPrintColor(ANSIForegroundRed, false))

	errorBefore := false

	// build got print
	for i := 0; i < min(expectedLen, gotLen); i++ {
		// same
		if expectedBytes[i] == gotBytes[i] {
			// remove error highlight
			if errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, false))
				errorBefore = false
			}

			// diff
		} else {
			// add error highlight
			if !errorBefore {
				print.Write(DiffPrintColor(ANSIForegroundRed, true))
				errorBefore = true
			}
		}

		print.WriteByte(gotBytes[i])
	}

	// add got trailing bytes
	if expectedLen < gotLen {
		if !errorBefore {
			print.Write(DiffPrintColor(ANSIForegroundRed, true))
			errorBefore = true
		}

		print.Write(gotBytes[expectedLen:])
	}

	// reset ANSI styles
	print.Write(ColorReset)
	print.WriteByte('\n')

	if opts.printToASCII {
		return errors.New(strconv.QuoteToASCII(print.String()))
	} else {
		return errors.New(print.String())
	}
}
