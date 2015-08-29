// Package scheduler is a local event scheduler responsible for managing
// simulation events (molecule diffusion, unimolecular decay, ...)
package scheduler

import (
	"github.com/haskelladdict/mcellLite/types"
)

// Event describes an event interface
type Event interface {
	Execute() error // execute the event
	Time() float64  // time of event
}

// Diffusion describes a diffusion event
type Diffusion struct {
	Mol *types.Molecule
	T   float64
}

// Execute method for Diffusion event
func (d Diffusion) Execute() error {
	// no op for now
	return nil
}

// Time method for diffusion event
func (d Diffusion) Time() float64 {
	return d.T
}
