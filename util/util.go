// Package util contains a collection of useful utility functions
package util

import "math"

func Equal(x, y float64) bool {
	return math.Nextafter(x, y)-y == 0.0
}
