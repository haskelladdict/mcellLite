// Package engine defines runtime functionality needed for reaction diffusion
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package engine

import (
	"fmt"
	"math"
	"math/rand"

	geom "github.com/haskelladdict/mcellLite/geometry"
	"github.com/haskelladdict/mcellLite/molecule"
	vec "github.com/haskelladdict/mcellLite/vector"
)

// Diffuse diffuses a molecule along the provided displacement vector,
// potentially reflecting off of mesh elements until the motion is complete.
func Diffuse(mol *molecule.Mol3, dt float64, mesh geom.Mesh, rng *rand.Rand) {
	// compute displacement
	scale := math.Sqrt(4 * mol.Spec.D * dt)
	disp := vec.V3{scale * rng.NormFloat64(), scale * rng.NormFloat64(),
		scale * rng.NormFloat64()}

	fmt.Println(disp, mol)
	hitPoint, ok := Collide(mol, disp, mesh)
	for ; ok; hitPoint, ok = Collide(mol, disp, mesh) {
		fmt.Println("diffusing ", disp)
		mol.MoveTo(hitPoint)
	}
	mol.MoveTo(vec.Add(mol.R, disp))
	fmt.Println(mol)
}

// Collide checks for collisions of molecule mol along disp with mesh elements
// If a collisions occurs it returns the hitPoint and true. If no collision
// occurs it returns nil and false
func Collide(mol *molecule.Mol3, disp vec.V3, mesh geom.Mesh) (vec.V3, bool) {

	for _, m := range mesh {
		hitPoint, status := geom.Intersect(mol.R, disp, &m)
		if status != 0 {
			continue // didn't hit mesh element
		}
		dispRem := vec.Sub(hitPoint, mol.R)

		// reflect: Rr = Ri - 2 N (Ri * N)
		disp = vec.Sub(dispRem, vec.Scalar(m.NN, 2*(vec.Dot(dispRem, m.NN))))

		// move slightly away from the triangle along the reflected ray.
		// If we happen to end our ray at hitpoint we move along the triangle
		// normal instead.
		if disp.Norm2() > geom.GEOM_EPSILON_2 {
			n := disp.Norm()
			dispN := vec.Scalar(disp, 1.0/n)
			hitPoint = vec.Add(hitPoint, vec.Scalar(dispN, geom.GEOM_EPSILON))
			disp = vec.Scalar(dispN, n-geom.GEOM_EPSILON)
		} else {
			side := 1.0
			if proj := vec.Dot(dispRem, m.NN); proj >= 0 {
				side = -1.0
			}
			hitPoint = vec.Add(hitPoint, vec.Scalar(m.NN, side*geom.GEOM_EPSILON))
		}
		return hitPoint, true
	}
	return vec.NullV3, false
}
