package conv

import (
	"strconv"
)

func StringToFloat64(buf string, dval float64) float64 {
	res, err := strconv.ParseFloat(buf, 64)

	if err == nil {
		return res
	}

	return dval
}

func Float64ToString(val float64) string {
	return strconv.FormatFloat(
		val,
		'f',
		-1,
		64,
	)
}

func StringToInt64(buf string, dval int64) int64 {
	res, err := strconv.ParseInt(buf, 10, 64)

	if err == nil {
		return res
	}

	return dval
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}
