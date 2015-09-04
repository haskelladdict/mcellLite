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
	"github.com/haskelladdict/mcellLite/vec"
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
	mol.MoveTo((mol.R).Add(disp))
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
		dispRem := hitPoint.Sub(mol.R)

		// reflect: Rr = Ri - 2 N (Ri * N)
		disp = dispRem.Sub((m.NN).Scalar(2 * (dispRem.Dot(m.NN))))

		// move slightly away from the triangle along the reflected ray.
		// If we happen to end our ray at hitpoint we move along the triangle
		// normal instead.
		if disp.Norm2() > geom.GEOM_EPSILON_2 {
			n := disp.Norm()
			dispN := disp.Scalar(1.0 / n)
			hitPoint = hitPoint.Add(dispN.Scalar(geom.GEOM_EPSILON))
			disp = dispN.Scalar(n - geom.GEOM_EPSILON)
		} else {
			side := 1.0
			if proj := dispRem.Dot(m.NN); proj >= 0 {
				side = -1.0
			}
			hitPoint = hitPoint.Add((m.NN).Scalar(side * geom.GEOM_EPSILON))
		}
		return hitPoint, true
	}
	return vec.NullV3, false
}
