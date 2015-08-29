// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package main

import (
	"fmt"

	geom "github.com/haskelladdict/mcellLite/geometry"
)

// Simulation contains all information related to simulation control such as
// timestep, number of iterations
// type Simulation struct {
// 	dt    float64
// 	iters uint64
// }

// // Partition contains all state information for a single patch in the simulation
// type Partition struct {
// 	Neighbors   []*Partition
// 	Mols        []*types.Molecule
// 	Faces       []*geometry.Face
// 	Events      scheduler.Events
// 	MolReceiver <-chan *types.Molecule
// 	Messager    chan struct{}
// 	Commander   chan int
// }

// makeCube returns an object corresponding to a regular cube
// func makeCube() geometry.Object {

// 	v0 := geometry.Vec3{1, 1, -1}
// 	v1 := geometry.Vec3{1, -1, -1}
// 	v2 := geometry.Vec3{-1, -1, -1}
// 	v3 := geometry.Vec3{-1, 1, 1}
// 	v4 := geometry.Vec3{1, 1, 1}
// 	v5 := geometry.Vec3{1, -1, 1}
// 	v6 := geometry.Vec3{-1, -1, 1}
// 	v7 := geometry.Vec3{-1, 1, 1}

// 	faces := []*geometry.Face{
// 		{v1, v2, v3},
// 		{v7, v6, v5},
// 		{v0, v4, v5},
// 		{v1, v5, v6},
// 		{v6, v7, v3},
// 		{v0, v3, v7},
// 		{v0, v1, v3},
// 		{v4, v7, v5},
// 		{v1, v0, v5},
// 		{v2, v1, v6},
// 		{v2, v6, v3},
// 		{v4, v0, v7},
// 	}

// 	return geometry.Object{faces}
// }

// Run is in charge of running all events occurring within the given
// partition. The main loop consists of checking for events on a cmd channel,
// several input channels or otherwise picking events of off the scheduler.
// func (p *Partition) Run() {
// 	// timer used for testing
// 	timer := time.NewTimer(time.Second * 2)

// 	// create scheduler
// 	sched := scheduler.Create()

// 	start := 0.0
// 	//	end := 1.0
// 	for {

// 		// add diffusion events
// 		for _, m := range p.Mols {
// 			sched.Add(&scheduler.Diffusion{m, start})
// 		}

// 		select {
// 		case c := <-p.Commander:
// 			fmt.Println("got command ", c)
// 		case m := <-p.MolReceiver:
// 			fmt.Println(m)
// 		case <-timer.C:
// 			fmt.Println("partition expiring")
// 			p.Messager <- struct{}{}
// 		}
// 	}
// }

func main() {
	// define vertices of geometry
	//  cube := makeCube()

	// // prepare partition
	// aMol := types.Species{6e-6}
	// mol1 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.5, 0.5, 0.6}}
	// mol2 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.3, 0.2, 0.4}}
	// mols := []*types.Molecule{&mol1, &mol2}

	// molCh := make(chan *types.Molecule)
	// cmdCh := make(chan int)
	// msgCh := make(chan struct{})
	// p := Partition{Neighbors: nil,
	// 	Mols:        mols,
	// 	Faces:       cube.Faces,
	// 	Events:      scheduler.Create(),
	// 	MolReceiver: molCh,
	// 	Messager:    msgCh,
	// 	Commander:   cmdCh,
	// }

	// go p.Run()
	// p.Commander <- 1
	// <-msgCh

	fmt.Println("Hi")
	m := geom.CreateRect(&geom.Vec3{-1.0, -1.0, -1.0}, &geom.Vec3{1.0, 1.0, 1.0})
	fmt.Println(m)
}

/*















*/
