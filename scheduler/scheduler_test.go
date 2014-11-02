// Package scheduler is a local event scheduler responsible for managing
// simulation events (molecule diffusion, unimolecular decay, ...)
package scheduler

import (
	"testing"

	"github.com/haskelladdict/mcellLite/geometry"
	"github.com/haskelladdict/mcellLite/types"
)

// Make sure the priority queue works as expected
func TestPriorityQueue(t *testing.T) {
	sched := Create()
	aMol := types.Species{6e-6}
	mol1 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.5, 0.5, 0.6}}
	mol2 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.3, 0.2, 0.4}}

	evt1 := Diffusion{&mol1, 0.5}
	evt2 := Diffusion{&mol2, 0.3}
	sched.Add(evt1)
	sched.Add(evt2)

	out1 := sched.Pop()
	out2 := sched.Pop()
	if out1.Time() < out2.Time() {
		t.Error("error in scheduler")
	}
}
