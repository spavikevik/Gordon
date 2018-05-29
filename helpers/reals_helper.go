package helpers

import (
	"strconv"
)

const (
	ε = 0.000000000000001
)

func Realf(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func Realeq(x, y float64) bool {
	return (x-y) < ε && (y-x) < ε
}
