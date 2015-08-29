// Package geometry contains geometry related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package geometry

import (
	"fmt"

	"github.com/haskelladdict/mcellLite/util"
)

// MeshElem describes a single triangular mesh element which may be part of
// a larger Mesh object.
type MeshElem struct {
	A, B, C Vec3 // triangle vertices
	U, V    Vec3 // u, v coordinates of mesh element
	N, NN   Vec3 // normal and normalized normal vectors
}

// NewMeshElem creates a new MeshElem and also properly computes the additional
// data members.
// NOTE: NewMeshElem assumes that v1, v2, v3 are not colinear, otherwise the
// function will panic
func NewMeshElem(v1, v2, v3 *Vec3) *MeshElem {
	u := v2.Sub(v1)
	v := v3.Sub(v1)
	n := u.Cross(v)

	// if the MeshElem is degenerate with panic
	if util.Equal(n.Norm(), 0.0) {
		panic(fmt.Sprintf("NewMeshElement: the provided vertices {%v, %v, %v} are colinear",
			v1, v2, v3))
	}
	return &MeshElem{A: *v1, B: *v2, C: *v3, U: *u, V: *v, N: *n, NN: *n.Scalar(1 / n.Norm())}
}

// Mesh is a collection of MeshElements
type Mesh []MeshElem

// CreateRect is a helper function creating a rectangular Mesh consisting of
// 12 individual MeshElem. The rectangle dimensions are specified by providing
// the coordinates of the lower left and upper right vertex
func CreateRect(llc, urc *Vec3) *Mesh {
	diag := urc.Sub(llc)
	c0 := llc
	c1 := llc.Add(&Vec3{diag.X, 0.0, 0.0})
	c2 := llc.Add(&Vec3{0.0, diag.Y, 0.0})
	c3 := llc.Add(&Vec3{0.0, 0.0, diag.Z})
	c4 := llc.Add(&Vec3{diag.X, diag.Y, 0.0})
	c5 := llc.Add(&Vec3{diag.X, 0.0, diag.Z})
	c6 := llc.Add(&Vec3{0.0, diag.Y, diag.Z})
	c7 := urc

	return &Mesh{
		*NewMeshElem(c0, c1, c5),
		*NewMeshElem(c0, c5, c3),
		*NewMeshElem(c1, c4, c7),
		*NewMeshElem(c1, c7, c5),
		*NewMeshElem(c4, c2, c6),
		*NewMeshElem(c4, c6, c7),
		*NewMeshElem(c2, c0, c3),
		*NewMeshElem(c2, c3, c6),
		*NewMeshElem(c5, c7, c6),
		*NewMeshElem(c5, c6, c3),
		*NewMeshElem(c0, c2, c1),
	}
}
