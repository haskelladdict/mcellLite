// Package geometry contains geometry related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package geometry

import (
	"math"

	"github.com/haskelladdict/mcellLite/util"
	vec "github.com/haskelladdict/mcellLite/vector"
)

// these epsilon is used for geometrical comparison. Anything smaller than that
// is assumed to be identical.
// FIXME: This is currently chosen arbitrarily and requires more thinking.
const GEOM_EPSILON = 1e-12
const GEOM_EPSILON_2 = 1e-24

// Intersect tests for ray triangle intersections. Possible return values are
//  0: triangle and ray segment intersect, in this case hitPoint contains the
//     location of the intersection point
//  1: triangle and ray intersect but ray segment does not reach the triangle
//  2: triangle and ray do not intersect
//  3: ray and triangle are co-planar
//  4: triangle is degenerate
//
// NOTE: This function was adapted from Dan Sunday
// <http://geomalgorithms.com/a06-_intersect-2.html#intersect3D_RayTriangle()>
func Intersect(start, disp vec.V3, m *MeshElem) (vec.V3, int) {

	// if the normal vector is zero, triangle is degenerate
	if vec.Equal(m.N, vec.NullV3) {
		return vec.NullV3, 4
	}

	// compute intersection of ray from p0 along disp with plane in which m is
	// located
	w0 := vec.Sub(start, m.A)
	a := vec.Dot(m.N, w0)
	b := vec.Dot(m.N, disp)
	if math.Abs(b) < GEOM_EPSILON { // our ray is parallel to triangle plane
		if util.Equal(a, 0.0) { // our ray is coplanar with the triangle
			return vec.NullV3, 3
		} else {
			return vec.NullV3, 2
		}
	}

	r := a / b
	if r < 0 { // if ray points away from triangle plane we won't hit it
		return vec.NullV3, 2
	} else if r > 1 { // if the ray segment doesn't reach the plane we won't hit it
		return vec.NullV3, 1
	}
	hitPoint := vec.Add(start, vec.Scalar(disp, r))

	// now test that hitPoint is within the triangle
	// we use local variables for efficiency
	w := vec.Sub(hitPoint, m.A)
	uu := vec.Dot(m.U, m.U)
	uv := vec.Dot(m.U, m.V)
	vv := vec.Dot(m.V, m.V)
	wu := vec.Dot(w, m.U)
	wv := vec.Dot(w, m.V)
	D := uv*uv - uu*vv

	// compute and test parametric coords
	s := (uv*wv - vv*wu) / D
	if s < 0.0 || s > 1.0 { // hitPoint is outside m
		return hitPoint, 2
	}
	t := (uv*wu - uu*wv) / D
	if t < 0.0 || (s+t) > 1.0 { // hitPoint is outside m
		return hitPoint, 2
	}
	return hitPoint, 0 // hitPoint is in m
}
