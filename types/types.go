// Package types defines the different MCell types
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package types

import (
	"github.com/haskelladdict/mcellLite/geometry"
)

// Species represents a molecule type
type Species struct {
	Name string  // species name
	D    float64 // diffusion coefficient
}

// Molecule represents a volume (3D) molecule
type Mol3 struct {
	Spec *Species      // what species are we
	R    geometry.Vec3 // where are we
	Bday float64       // when were we born
}
