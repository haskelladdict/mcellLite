// Package vec contains vector related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package vec

import (
	"math"

	"github.com/haskelladdict/mcellLite/util"
)

// Vec3 is a 3D vector
type Vec3 struct {
	X, Y, Z float64
}

// NullVec3 is a convenience definition for the null vector
var NullVec3 = Vec3{0.0, 0.0, 0.0}

// Norm2 is a method for Vec3 computing its squared L2 norm
func (v Vec3) Norm2() float64 {
	return v.Dot(v)
}

// Norm is a method for Vec3 computing its L2 norm
func (v Vec3) Norm() float64 {
	return math.Sqrt(v.Norm2())
}

// Scalar computes the scalar multiplication of scalar a with Vec3 v
func (v Vec3) Scalar(a float64) Vec3 {
	return Vec3{a * v.X, a * v.Y, a * v.Z}
}

// Sub subtracts Vec3 v and Vec3 w
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Add adds Vec3 v and Vec3 w
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Cross computes the cross product of Vec3 v with Vec3 w
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// Equal determines if Vec3 v and w are identical. Identity is determined via
// component-wise identity
func (v Vec3) Equal(w Vec3) bool {
	return util.Equal(v.X, w.X) && util.Equal(v.Y, w.Y) && util.Equal(v.Z, w.Z)
}

// Dot computes the dot product of Vec3 v with Vec3 w
func (v Vec3) Dot(w Vec3) float64 {
	return (v.X * w.X) + (v.Y * w.Y) + (v.Z * w.Z)
}
