// Package geometry contains geometry related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package geometry

import (
	"testing"

	"github.com/haskelladdict/mcellLite/util"
)

// Make sure Vec3 works as expected
func TestVec3(t *testing.T) {

	v1 := &Vec3{1.0, 0.0, 0.0}
	if !util.Equal(v1.Norm2(), 1.0) || !util.Equal(v1.Norm(), 1.0) {
		t.Errorf("Norm: expected 1.0, got %15.15f; Norm2: expected 1.0, got %15.15f ",
			v1.Norm(), v1.Norm2())
	}

	v1 = &Vec3{1.0, 1.0, 1.0}
	if !util.Equal(v1.Norm2(), 3.0) || !util.Equal(v1.Norm(), 1.732050807568877) {
		t.Errorf("Norm: expected 1.73205080756888, got %15.15f; Norm2: expected 3.0, got %15.15f ",
			v1.Norm(), v1.Norm2())
	}

	v2 := v1.Add(v1)
	if !v2.Equal(&Vec3{2.0, 2.0, 2.0}) {
		t.Errorf("Vector addition: expected {2.0, 2.0, 2.0} got %v", v2)
	}

	v2 = v1.Sub(v1)
	if !v2.Equal(&Vec3{0.0, 0.0, 0.0}) {
		t.Errorf("Vector subtraction: expected {0.0, 0.0, 0.0} got %v", v2)
	}

	v2 = v1.Scalar(5.0)
	if !util.Equal(v2.X, 5.0) || !util.Equal(v2.Y, 5.0) || !util.Equal(v2.Z, 5.0) {
		t.Errorf("Scalar multiplication: expected {5.0, 5.0, 5.0} got %v", v2)
	}

	a := v1.Dot(v2)
	if !util.Equal(a, 15.0) {
		t.Errorf("Dot product: expected 15 got %15.15f", a)
	}

	v1 = &Vec3{1.0, 0.0, 0.0}
	v2 = &Vec3{0.0, 1.0, 0.0}
	v3 := v1.Cross(v2)
	if !v3.Equal(&Vec3{0.0, 0.0, 1.0}) {
		t.Errorf("Cross product: expected {0.0, 0.0, 1.0} got %v", v3)
	}
}
