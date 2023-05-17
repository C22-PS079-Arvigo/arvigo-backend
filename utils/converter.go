package utils

import (
	"math"
	"strconv"
	"strings"
)

func StrToInt(text string, defaultReturn int) int {
	number := defaultReturn
	if text != "" {
		var err error
		number, err = strconv.Atoi(text)
		if err != nil {
			number = defaultReturn
		}
	}
	return number
}

func StrToInt64(text string, defaultReturn int64) int64 {
	number := defaultReturn
	if text != "" {
		var err error
		number, err = strconv.ParseInt(text, 10, 64)
		if err != nil {
			number = defaultReturn
		}
	}
	return number
}

func StrToUint64(text string, defaultReturn uint64) uint64 {
	number := defaultReturn
	if text != "" {
		var err error
		number, err = strconv.ParseUint(text, 10, 64)
		if err != nil {
			number = defaultReturn
		}
	}
	return number
}

// This function is used to round the fraction value after the comma to a certain number of digits
func RoundFloat64(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// helper function for compare float
const float64EqualityThreshold = 0.1

// AlmostEqual return true if two float diferent < float64EqualityThreshold
func AlmostEqual(a, b float64) bool {
	dif := math.Abs(a - b)
	if dif < 0 {
		dif = -dif
	}
	return dif <= float64EqualityThreshold
}

func Uint32ToString(val uint32) string {
	return strconv.FormatUint(
		uint64(val),
		10,
	)
}

func StringToInt64(buf string, dval int64) int64 {
	res, err := strconv.ParseInt(buf, 10, 64)

	if err == nil {
		return res
	}

	return dval
}

func IntToString(val int) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}

func StringToInt(buf string, dval int) int {
	res, err := strconv.ParseInt(buf, 10, 0)

	if err == nil {
		return int(res)
	}

	return dval
}

func StringToUint8(buf string, dval uint8) uint8 {
	res, err := strconv.ParseUint(buf, 10, 8)

	if err == nil {
		return uint8(res)
	}

	return dval
}

func StringToUint32(buf string, dval uint32) uint32 {
	res, err := strconv.ParseUint(buf, 10, 32)

	if err == nil {
		return uint32(res)
	}

	return dval
}

func IntToPtr(number int) *int {
	return &number
}

func Float64ToPtr(number float64) *float64 {
	return &number
}

func BoolToInt(bol bool) int {
	if bol {
		return 1
	}
	return 0
}

func ConvertStringToSlice(str, delimiter string) (res []string) {
	for _, val := range strings.Split(str, delimiter) {
		if strings.TrimSpace(val) != "" {
			res = append(res, val)
		}
	}

	return
}
