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
	D float64
}

// Molecule represents a generic molecule
type Molecule struct {
	Species               // what species are we
	Pos     geometry.Vec3 // where are we
}
