// Package vec contains vector related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package vec

import (
	"math"

	"github.com/haskelladdict/mcellLite/util"
)

// V3 is a 3D vector
type V3 struct {
	X, Y, Z float64
}

// NullV3 is a convenience definition for the null vector
var NullV3 = V3{0.0, 0.0, 0.0}

// Norm2 is a method for V3 computing its squared L2 norm
func (v V3) Norm2() float64 {
	return v.Dot(v)
}

// Norm is a method for V3 computing its L2 norm
func (v V3) Norm() float64 {
	return math.Sqrt(v.Norm2())
}

// Scalar computes the scalar multiplication of scalar a with V3 v
func (v V3) Scalar(a float64) V3 {
	return V3{a * v.X, a * v.Y, a * v.Z}
}

// Sub subtracts V3 v and V3 w
func (v V3) Sub(w V3) V3 {
	return V3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Add adds V3 v and V3 w
func (v V3) Add(w V3) V3 {
	return V3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Cross computes the cross product of V3 v with V3 w
func (v V3) Cross(w V3) V3 {
	return V3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// Equal determines if V3 v and w are identical. Identity is determined via
// component-wise identity
func (v V3) Equal(w V3) bool {
	return util.Equal(v.X, w.X) && util.Equal(v.Y, w.Y) && util.Equal(v.Z, w.Z)
}

// Dot computes the dot product of V3 v with V3 w
func (v V3) Dot(w V3) float64 {
	return (v.X * w.X) + (v.Y * w.Y) + (v.Z * w.Z)
}
