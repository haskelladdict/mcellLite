// Package geometry contains geometry related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package geometry

import (
	"math"

	"github.com/haskelladdict/mcellLite/util"
)

// these epsilon is used for geometrical comparison. Anything smaller than that
// is assumed to be identical.
// FIXME: This is currently chosen arbitrarily and requires more thinking.
const geom_EPSILON = 1e-12
const geom_EPSILON_2 = 1e-24

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
func Intersect(p0, disp *Vec3, m *MeshElem) (*Vec3, int) {

	// if the normal vector is zero, triangle is degenerate
	if m.N.Equal(&Vec3{0.0, 0.0, 0.0}) {
		return nil, 4
	}

	// compute intersection of ray from p0 along disp with plane in which m is
	// located
	w0 := p0.Sub(&m.A)
	a := -m.N.Dot(w0)
	b := m.N.Dot(disp)
	if math.Abs(b) < geom_EPSILON { // our ray is parallel to triangle plane
		if util.Equal(a, 0.0) { // our ray is coplanar with the triangle
			return nil, 3
		} else {
			return nil, 2
		}
	}

	r := a / b
	if r < 0 { // if ray points away from triangle plane we won't hit it
		return nil, 2
	} else if r > 1 { // if the ray segment doesn't reach the plane we won't hit it
		return nil, 1
	}
	hitPoint := p0.Add(disp.Scalar(r))

	// now test that hitPoint is within the triangle
	// we use local variable for efficiency
	w := hitPoint.Sub(&m.A)
	u := &m.U
	v := &m.V
	uu := u.Dot(u)
	uv := u.Dot(v)
	vv := v.Dot(v)
	wu := w.Dot(u)
	wv := w.Dot(v)
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
