// Package mol defines the different MCell molecule types
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package mol

import "github.com/haskelladdict/mcellLite/vec"

// MolMap is a dictionary with species names as keys and an arrays of molecules
// of that type a values
type MolMap map[string][]*Mol3

// Species represents a molecule type
type Species struct {
	Name string  // species name
	D    float64 // diffusion coefficient
}

// Molecule represents a volume (3D) molecule
type Mol3 struct {
	Spec *Species // what species are we
	R    vec.Vec3 // where are we
	Bday float64  // when were we born
}

// moveTo moves the given molecule to the specified position
func (m *Mol3) MoveTo(newPos vec.Vec3) {
	m.R = newPos
}
