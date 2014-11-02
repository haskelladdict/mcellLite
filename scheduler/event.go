// Package scheduler is a local event scheduler responsible for managing
// simulation events (molecule diffusion, unimolecular decay, ...)
package scheduler

import (
	"github.com/haskelladdict/mcellLite/types"
)

// Event describes an event interface
type Event interface {
	Execute() error // execute the event
	Type() int      // type of event
	Time() float64  // time of event
}

// Diffusion describes a diffusion event
type Diffusion struct {
	mol *types.Molecule
	t   float64
}

// Execute method for Diffusion event
func (d Diffusion) Execute() error {
	// no op for now
	return nil
}

// Type method for Diffusion event
func (d Diffusion) Type() int {
	return 0
}

// Time method for diffusion event
func (d Diffusion) Time() float64 {
	return d.t
}
